# need: editdistance, prestring(go branch)
import editdistance
import sys
import contextlib
import argparse
import copy
import json
import logging
from collections import ChainMap, deque, defaultdict, namedtuple
from prestring.go import GoModule

Action = namedtuple("Action", "action, src, dst")
logger = logging.getLogger(__name__)


class Reader(object):
    def read_world(self, data, parent=None):
        world = World(parent=parent, reader=self)
        for name, module in data["module"].items():
            world.read_module(name, module)
        return world

    def read_module(self, data, parent=None):
        module = Module(data["name"], data["fullname"], parent=parent, reader=self)
        for name, file in data["file"].items():
            module.read_file(name, file)
        return module

    def read_file(self, data, parent=None):
        file = File(data["name"], data["import"], parent=parent, reader=self)
        for name, alias in data["alias"].items():
            file.read_alias(name, alias)
        for name, struct in data["struct"].items():
            file.read_struct(name, struct)
        return file

    def read_struct(self, data, parent=None):
        struct = Struct(data["name"], data, parent=parent, reader=self)
        return struct

    def read_alias(self, data, parent=None):
        alias = Alias(data["name"], data, parent=parent, reader=self)
        return alias


class World(object):
    def __init__(self, parent=None, reader=None):
        self.parent = parent
        self.reader = reader
        self.modules = {}
        self.modules_by_fullname = {}

    def read_module(self, name, module):
        self.modules[name] = self.reader.read_module(module, parent=self)
        self.modules_by_fullname[self.modules[name].fullname] = self.modules[name]

    def normalize(self):
        for module in self.modules.values():
            module.normalize()

    def __getitem__(self, name):
        return self.modules[name]


class Module(object):
    def __init__(self, name, fullname, parent=None, reader=None):
        self.name = name
        self.fullname = fullname
        self.parent = parent
        self.reader = reader
        self.files = {}
        self.reset()

    @property
    def package_name(self):
        return self.name

    def with_fullname(self, name):
        return "{}.{}".format(self.fullname, name)

    def read_file(self, name, file):
        self.files[name] = self.reader.read_file(file, parent=self)

    def normalize(self):
        self.reset()
        for file in self.files.values():
            file.normalize()
            self.new_child(file.members)

    def reset(self):
        self.members = ChainMap()

    def new_child(self, item):
        self.members = self.members.new_child(item)

    def __getitem__(self, name):
        return self.members[name]

    def get(self, name, default=None):
        return self.members.get(name, default)

    def __contains__(self, name):
        return name in self.members


class File(object):
    def __init__(self, name, imports, parent=None, reader=None):
        self.name = name
        self.imports = imports
        self.parent = parent
        self.reader = reader
        self.aliases = {}
        self.structs = {}
        self.members = {}

    @property
    def package_name(self):
        if self.parent is None:
            return None
        return self.parent.package_name

    def normalize(self):
        for struct in self.structs.values():
            struct.normalize()

        for alias in self.aliases.values():
            alias.normalize()

    def read_alias(self, name, alias):
        self.members[name] = self.aliases[name] = self.reader.read_alias(alias, parent=self)

    def read_struct(self, name, struct):
        self.members[name] = self.structs[name] = self.reader.read_struct(struct, parent=self)

    def dump(self, writer):
        return writer.write_file(self)


class Alias(object):
    def __init__(self, name, data, parent=None, reader=None):
        self.name = name
        self.parent = parent
        self.reader = reader
        self.data = data

    @property
    def package_name(self):
        if self.parent is None:
            return None
        return self.parent.package_name

    def dump(self, writer):
        return writer.write_alias(self)

    def normalize(self):
        pass


class Struct(object):
    def __init__(self, name, data, parent=None, reader=None):
        self.name = name
        self.parent = parent
        self.reader = reader
        self.rawdata = data
        self.data = data

    @property
    def package_name(self):
        if self.parent is None:
            return None
        return self.parent.package_name

    @property
    def fullname(self):
        return self.parent.parent.with_fullname(self.name)

    def dump(self, writer):
        return writer.write_struct(self)

    def __getitem__(self, name):
        return self.data["fields"][name]

    def __contains__(self, name):
        return name in self.data["fields"]

    def fields(self):
        # todo: field class
        return self.data["fields"].items()

    def normalize(self):
        self.rawdata = copy.deepcopy(self.rawdata)
        data = self.data["fields"]
        for k in list(data.keys()):
            data[k.lower()] = data.pop(k)


class ConvertWriter(object):
    def __init__(self, m, convertor):
        self.m = m
        self.convertor = convertor

    def get_function_name(self, src, dst):
        return 'ConvertFrom{}{}'.format(src.package_name.title(), dst.name)

    def write(self, src, dst):
        fnname = self.get_function_name(src, dst)
        src_type_list = tuple(["pointer", src.fullname])
        dst_type_list = tuple(["pointer", dst.fullname])
        cont = []
        self._register(fnname, src, dst, src_type_list, dst_type_list)
        self._write(fnname, src, dst, cont)
        for fn in cont:
            fn()

    def _register(self, fnname, src, dst, src_type_list, dst_type_list):
        override = self.convertor.as_override

        @override(src_type_list, dst_type_list)
        def subconvert(convertor, m, value, *args):
            tmp = convertor.tmp_name()
            m.stmt("{}, err := {}({})".format(tmp, fnname, value))
            with m.if_("err != nil"):
                m.return_("nil, err")
            return tmp

    def _write(self, fnname, src, dst, cont):
        src_arg = 'src *{}.{}'.format(src.package_name, src.name)
        dst_arg = '*{}.{}'.format(dst.package_name, dst.name)
        with self.m.func(fnname, src_arg, return_="({}, error)".format(dst_arg)):
            self.m.stmt("dst := &{}.{}{{}}".format(dst.package_name, dst.name))
            for name, field in sorted(dst.fields()):
                self.write_code_convert(src, dst, name, field, cont)
            self.m.return_("dst, nil")

    def write_code_convert(self, src, dst, name, field, cont, retry=False):
        try:
            if name in src:
                value = "src.{}".format(src[name]["name"])
                convert = self.convertor.convert(self.m, src[name], field, src, dst, name=value)
                self.m.stmt("dst.{} = {}".format(field["name"], convert))  # xxx
            else:
                try:
                    score, nearlest = min([(editdistance.eval(name, f["name"]), f["name"]) for _, f in src.fields()])
                    self.m.comment("FIXME: {} is not found. (maybe {}?)".format(field["name"], nearlest))
                except ValueError:
                    self.m.comment("FIXME: {} is not found.".format(field["name"]))
        except GencodeMappingNotFound as e:
            # print("@", e, file=sys.stderr)
            if retry:
                raise

            if e.src[-1].endswith(("32", "64")) or e.dst[-1].endswith(("32", "64")):
                primitive_src = e.src[-1].replace("32", "").replace("64", "")
                primitive_dst = e.src[-1].replace("32", "").replace("64", "")
                if primitive_src == primitive_dst:
                    self.convertor.codegen.resolver.add_relation(e.src[-1], e.dst[-1])
                    return self.write_code_convert(src, dst, name, field, cont, retry=True)

            # xxx:
            dep_src_type_list = tuple(get_type_path(src[name]["type"], src))
            dep_dst_type_list = tuple(get_type_path(field["type"], dst))
            src_prefix, src_name = dep_src_type_list[-1].rsplit(".", 1)
            dst_prefix, dst_name = dep_dst_type_list[-1].rsplit(".", 1)
            dep_src = src.parent.parent.parent.modules_by_fullname[src_prefix][src_name]
            dep_dst = dst.parent.parent.parent.modules_by_fullname[dst_prefix][dst_name]
            fnname = self.get_function_name(dep_src, dep_dst)
            self._register(fnname, dep_src, dep_dst, dep_src_type_list, dep_dst_type_list)
            cont.append(lambda: self._write(fnname, dep_src, dep_dst, cont))
            return self.write_code_convert(src, dst, name, field, cont, retry=True)


class ImportWriter(object):
    def __init__(self, im):
        self.im = im
        self.prefix_map = {}  # fullname -> prefix
        self.name_map = {}  # name -> fullname
        self.used = set()
        self.i = 0

    def _get_prefix(self, module):
        fullname = self.name_map.get(module.name)
        if fullname is None:
            self.name_map[module.name] = module.fullname
            return module.name
        elif fullname == module.fullname:
            return module.name
        else:
            prefix = self.prefix_map.get(module.fullname)
            if prefix is None:
                prefix = self.prefix_map[module.fullname] = "{}{}".format(module.name, self.i)
                self.i += 1
            return prefix

    def import_(self, module):
        prefix = self._get_prefix(module)
        if module.fullname in self.used:
            return prefix
        self.used.add(module.fullname)
        self.im.import_(module.fullname, as_=prefix)
        return prefix


def get_type_path(value, struct):
    if value["kind"] == "primitive":
        module = struct.parent.parent
        if value["value"] not in module:
            return [value["value"]]
        else:
            return ["{}.{}".format(module.fullname, value["value"])]
    elif value["kind"] == "selector":
        prefix = struct.parent.imports[value["prefix"]]["fullname"]
        return ["{}.{}".format(prefix, value["value"])]
    else:
        r = [value["kind"]]
        r.extend(get_type_path(value["value"], struct))
        return r


class TypeConvertor(object):
    def __init__(self, import_writer, codegen, src_world, dst_world):
        self.import_writer = import_writer
        self.codegen = codegen
        self.src_world = src_world
        self.dst_world = dst_world
        self.i = 0
        items = []
        for world, other_world in [(src_world, dst_world), (dst_world, src_world)]:
            for _, module in world.modules.items():
                for _, file in module.files.items():
                    for _, alias in file.aliases.items():
                        data = alias.data
                        src_name = data["original"]["value"]
                        if "prefix" in data["original"]:
                            # xxx:
                            src_fullname = "{}.{}".format(file.imports[data["original"]["prefix"]]["fullname"], src_name)
                            items.append((src_fullname, module.with_fullname(data["name"])))
                        else:
                            items.append((src_name, module.with_fullname(data["name"])))
        self.codegen.resolver.add_relation_list(items)
        self.override_map = {}

    def as_override(self, src_type, dst_type):
        def decorator(on_write):
            self.override(src_type, dst_type, on_write)
            return on_write
        return decorator

    def override(self, src_type, dst_type, on_write):
        self.codegen.resolver.add_relation(src_type, dst_type)
        self.override_map[(src_type, dst_type)] = on_write

    def get_new_prefix_on_selector(self, prefix, src_struct, dst_struct):
        dst_module = dst_struct.parent.parent
        if dst_module.name == prefix:
            fullname = dst_module.fullname
        else:
            fullname = dst_struct.parent.imports[prefix]["fullname"]
        module = dst_module.parent.modules_by_fullname[fullname]
        # writing import clause if needed
        return self.import_writer.import_(module)

    def coerce(self, m, value, src_type, dst_type, src_struct, dst_struct):
        pair = (src_type, dst_type)
        if pair in self.override_map:
            return self.override_map[pair](self, m, value, src_type, dst_type, src_struct, dst_struct)

        if isinstance(dst_type, (list, tuple)):
            return "({})({})".format("".join("*" if x == "pointer" else x for x in dst_type), value)
        elif "/" in dst_type:
            prefix_and_name = dst_type.rsplit("/", 1)[-1]
            prefix, name = prefix_and_name.rsplit(".")
            new_prefix = self.get_new_prefix_on_selector(prefix, src_struct, dst_struct)
            return "{}.{}({})".format(new_prefix, name, value)
        else:
            return "{}({})".format(dst_type, value)

    def tmp_name(self):
        self.i += 1
        return "tmp{}".format(self.i)

    def convert(self, m, src, dst, src_struct, dst_struct, name):
        src_path = list(get_type_path(src["type"], src_struct))
        dst_path = list(get_type_path(dst["type"], dst_struct))
        code = self.codegen.gencode(src_path, dst_path)
        value = name
        is_cast = False
        # print("##", code, file=sys.stderr)
        for op in code:
            if is_cast:
                tmp = self.tmp_name()
                m.stmt("{} := {}".format(tmp, value))
                value = tmp
                is_cast = False

            if op[0] == "deref":
                value = "*({})".format(value)
            elif op[0] == "ref":
                value = "&({})".format(value)
                is_cast = True
            elif op[0] == "coerce":
                value = self.coerce(m, value, op[1], op[2], src_struct, dst_struct)
                is_cast = True
            else:
                m.comment("hmm {}".format(op[0]))
                # raise Exception(op[0])
        return value


class TypeMappingResolver(object):
    def __init__(self, items):
        self.primitive_map = defaultdict(list)
        self.detail_map = defaultdict(list)
        self.add_relation_list(items)

    def add_relation_list(self, relations):
        for k, v in relations:
            self.add_relation(k, v)

    def add_relation(self, k, v):
        k = self._wrap_value(k)
        v = self._wrap_value(v)
        self._add_value(k, v)
        self._add_value(v, k)

    def _has_detail(self, wv):
        return len(wv) > 1

    def _add_value(self, k, v):
        self.detail_map[k].append(v)
        if self._has_detail(k):
            self.primitive_map[k[-1:]].append((k, v))
            self.primitive_map[v[-1:]].append((v, k))

    def _wrap_value(self, v):
        if isinstance(v, (tuple, list)):
            return tuple(v)
        else:
            return (v,)

    def _unwrap_value(self, v):
        if not self._has_detail(v):
            return v[0]
        else:
            return v

    def resolve(self, src, dst):
        if src == dst:
            return []
        src = self._wrap_value(src)
        dst = self._wrap_value(dst)
        if src[-1:] == dst[-1:]:
            return self.on_finish([src], [dst])
        else:
            # print("detail_map=", self.detail_map)
            # print("primitive_map=", self.primitive_map)
            return self.tick_src([src], [dst], set(), set(), deque(), deque())

    def on_finish(self, src_hist, dst_hist):
        src_hist.extend(reversed(dst_hist))
        path = src_hist
        coerce_path = []
        for i in range(len(path) - 1):
            if path[i] == path[i + 1]:
                continue
            prev, next_ = self._unwrap_value(path[i]), self._unwrap_value(path[i + 1])
            if coerce_path and coerce_path[-1][1] == next_ and coerce_path[-1][2] == prev:
                coerce_path.pop()
                continue
            coerce_path.append(Action(action="coerce", src=prev, dst=next_))
        return coerce_path

    def tick_src(self, src_hist, dst_hist, src_arrived, dst_arrived, src_q, dst_q):
        # print("@S", "src_hist=", src_hist, "dst_hist=", dst_hist, "src_arrived=", src_arrived, "src_q=", src_q, "dst_q=", dst_q)
        if src_hist[-1][-1] == dst_hist[-1][-1]:
            if src_hist[-1] == dst_hist[-1]:
                return self.on_finish(src_hist, dst_hist[:-1])
            else:
                return self.on_finish(src_hist, dst_hist)
        if src_hist[-1] not in src_arrived:
            src_arrived.add(src_hist[-1])
            if src_hist[-1] in self.detail_map:
                for item in self.detail_map[src_hist[-1]]:
                    src_q.append(([*src_hist, item], dst_hist))

            for i in range(1, len(src_hist[-1])):
                k = src_hist[-1][i:]
                if k in self.detail_map:
                    for item in self.detail_map[k]:
                        src_q.append(([*src_hist, k, item], dst_hist))
            k = src_hist[-1][-1:]
            if k in self.primitive_map:
                for items in self.primitive_map[k]:
                    src_q.append(([*src_hist, *items], dst_hist))
        if dst_q:
            src_hist, dst_hist = dst_q.pop()
            return self.tick_dst(src_hist, dst_hist, src_arrived, dst_arrived, src_q, dst_q)
        elif src_q:
            return self.tick_dst(src_hist, dst_hist, src_arrived, dst_arrived, src_q, dst_q)
        else:
            return None

    def tick_dst(self, src_hist, dst_hist, src_arrived, dst_arrived, src_q, dst_q):
        # print("@D", "src_hist=", src_hist, "dst_hist=", dst_hist, "src_arrived=", src_arrived, "src_q=", src_q, "dst_q=", dst_q)
        if src_hist[-1][-1] == dst_hist[-1][-1]:
            if src_hist[-1] == dst_hist[-1]:
                return self.on_finish(src_hist, dst_hist[:-1])
            else:
                return self.on_finish(src_hist, dst_hist)
        if dst_hist[-1] not in dst_arrived:
            dst_arrived.add(dst_hist[-1])
            if dst_hist[-1] in self.detail_map:
                for item in self.detail_map[dst_hist[-1]]:
                    dst_q.append((src_hist, [*dst_hist, item]))

            for i in range(1, len(dst_hist[-1])):
                k = dst_hist[-1][i:]
                if k in self.detail_map:
                    for item in self.detail_map[k]:
                        dst_q.append((src_hist, [*dst_hist, k, item]))
            k = dst_hist[-1][-1:]
            if k in self.primitive_map:
                for items in self.primitive_map[k]:
                    dst_q.append((src_hist, [*dst_hist, *items]))
            return self.tick_src(src_hist, dst_hist, src_arrived, dst_arrived, src_q, dst_q)
        if src_q:
            src_hist, dst_hist = src_q.pop()
            return self.tick_src(src_hist, dst_hist, src_arrived, dst_arrived, src_q, dst_q)
        elif dst_q:
            return self.tick_src(src_hist, dst_hist, src_arrived, dst_arrived, src_q, dst_q)
        else:
            return None


class GencodeMappingNotFound(ValueError):
    def __init__(self, msg, src, dst):
        super().__init__(msg)
        self.src = src
        self.dst = dst


class MiniCodeGenerator(object):
    # generating mini language
    # [deref] -> *x
    # [ref] -> &x
    # [coerce, x, y] -> finding alias {original: x} and name is y, -> y(x)

    def __init__(self, resolver, optimizer=None):
        self.resolver = resolver
        self.optimizer = optimizer or self._default_optimize

    def _default_optimize(self, code):
        optimized = []
        for x in code:
            if x[0] == "ref" and optimized and optimized[-1][0] == "deref":
                optimized.pop()
            if x[0] == "iterate" and optimized and optimized[-1][0] == "deiterate":
                optimized.pop()
            else:
                optimized.append(x)
        return optimized

    def gencode(self, src_path, dst_path):
        code = self._gencode(src_path, dst_path)
        code = self.optimizer(code)
        return code

    def _gencode(self, src_path, dst_path):
        code = []
        mapping_path = self.resolver.resolve(src_path, dst_path)
        if mapping_path is None:
            msg = "mapping not found {!r} -> {!r}".format(src_path, dst_path)
            raise GencodeMappingNotFound(msg, src_path, dst_path)

        def get_primitive(v):
            if isinstance(v, (list, tuple)):
                return v[-1]
            else:
                return v

        for action in mapping_path:
            if get_primitive(action.src) == get_primitive(action.dst):
                if isinstance(action.src, (list, tuple)):
                    itr = reversed(action.src)
                    next(itr)
                    for typ in itr:
                        if typ == "pointer":
                            code.append(("deref", ))
                        elif typ == "array":
                            code.append(("deiterate", ))
                        else:
                            raise ValueError("not implemented: typ={}, path={}".format(typ, src_path[1:]))
                if isinstance(action.dst, (list, tuple)):
                    itr = reversed(action.dst)
                    next(itr)
                    for typ in itr:
                        if typ == "pointer":
                            code.append(("ref", ))
                        elif typ == "array":
                            code.append(("iterate", ))
                        else:
                            raise ValueError("not implemented: typ={}, path={}".format(typ, src_path[1:]))
            else:
                code.append(action)
        return code


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--src", required=True)
    parser.add_argument("--dst", required=True)
    args = parser.parse_args()

    reader = Reader()

    with open(args.src) as rf:
        src_world = reader.read_world(json.load(rf))
        src_world.normalize()
    with open(args.dst) as rf:
        dst_world = reader.read_world(json.load(rf))
        dst_world.normalize()
    gencoder = MiniCodeGenerator(TypeMappingResolver([]))

    m = GoModule()
    m.package("convert")
    with m.import_group() as im:
        iw = ImportWriter(im)
        convertor = TypeConvertor(iw, gencoder, src_world, dst_world)
        iw.import_(src_world["model"])
        iw.import_(dst_world["def"])

    @convertor.as_override("string", src_world["bson"].with_fullname("ObjectId"))
    def string_to_object_id(convertor, m, value, *args):
        module = src_world["bson"]
        new_prefix = convertor.import_writer.import_(module)
        return "{}.ObjectIdHex({})".format(new_prefix, value)

    @convertor.as_override(src_world["bson"].with_fullname("ObjectId"), "string")
    def object_id_to_string(convertor, m, value, *args):
        return "{}.Hex()".format(value)

    @convertor.as_override(src_world["model"].with_fullname("Date"), "time.Time")
    def model_date_to_time(convertor, m, value, *args):
        return "{}.Time()".format(value)

    m.sep()
    cw = ConvertWriter(m, convertor)
    for module in dst_world.modules.values():
        for file in module.files.values():
            for struct in file.structs.values():
                for module in src_world.modules.values():
                    if struct.name in module:
                        print("@", struct.name, file=sys.stderr)
                        cw.write(module[struct.name], struct)
                    elif struct.name.startswith("Enduser") and struct.name[len("Enduser"):] in module:
                        print("<", struct.name, file=sys.stderr)
                        cw.write(module[struct.name[len("Enduser"):]], struct)
                    elif struct.name.startswith("Tuner") and struct.name[len("Tuner"):] in module:
                        print(">", struct.name, file=sys.stderr)
                        cw.write(module[struct.name[len("Tuner"):]], struct)

    # cw.write(src_world["model"]["Page"], dst_world["def"]["Page"])
    # # cw.write(dst_world["def"]["Page"], src_world["model"]["Page"])
    # cw.write(src_world["model"]["User"], dst_world["def"]["User"])

    print(m)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    main()
