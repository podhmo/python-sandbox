import typing as t
from functools import singledispatchmethod
from collections import defaultdict

# todo: nullable
# todo: hetero list
# todo: many=True
# todo: zero

uid = str
TypeInfo = t.Union["Object", "Container", "Primitive"]
Path = t.List[str]
JSONType = t.Union[int, float, str, bool, t.List["JSONType"], t.Dict[str, "JSONType"]]


class Object:
    def __init__(
        self, sig: t.Any, *, path: Path, props: t.Dict, raw: t.Dict[t.Any, t.Any]
    ) -> None:
        self.sig = sig
        self.path = path
        self.props = props
        self._raw = raw
        self._others: t.List[t.Tuple[Path], t.Dict[t.Any, t.Any]] = []

    @property
    def size(self):
        return len(self.props)

    def __repr__(self):
        return f"<Object size={self.size} path={self.path!r}>"


class Container:
    def __init__(
        self,
        sig: t.Any,
        *,
        path: Path,
        base: t.Type[t.Any],
        item: TypeInfo,
        raw: t.Dict[t.Any, t.Any],
    ) -> None:
        self.sig = sig
        self.path = path
        self.base = base
        self.item = item
        self._raw = raw
        self._others: t.List[t.Tuple[Path], t.Dict[t.Any, t.Any]] = []

    @property
    def size(self):
        return len(self._raw)

    def __repr__(self):
        return f"<Container base={self.base} size={self.size} path={self.path!r}>"


class Primitive:
    def __init__(
        self,
        sig: t.Any,
        *,
        path: Path,
        type_: t.Type[t.Any],
        raw: t.Dict[t.Any, t.Any],
    ) -> None:
        self.sig = sig
        self.path = path
        self.type = type_
        self._raw = raw

    @property
    def size(self):
        return 0

    def __repr__(self):
        return f"<Primitive type={self.type!r} path={self.path!r}>"


SENTINEL = object()
ZERO = Object(tuple(), path=":zero:", props={}, raw={})


class Result:
    def __init__(self):
        self.registry: t.Dict[uid, TypeInfo] = {}
        self.history: t.List[TypeInfo] = []
        self._uid_map: t.Dict[t.Any, uid] = defaultdict(lambda: len(self._uid_map))

    def __repr__(self):
        return f"<Result size={len(self.history)}>"

    def add(self, uid: uid, info: TypeInfo) -> TypeInfo:
        self.registry[uid] = info
        self.history.append(info)  # xxx
        return info

    def __contains__(self, uid: uid):
        return uid in self.registry

    def get(self, uid: uid, *, default):
        return self.registry.get(uid, default)


class Detector:
    # todo: many
    def detect(self, d: JSONType, *, path: Path, result: Result) -> TypeInfo:
        if isinstance(d, dict):
            return self.detect_dict(d, path=path, result=result)
        elif isinstance(d, (list, tuple)):
            return self.detect_list(d, path=path, result=result)
        else:
            return self.detect_primitive(d, path=path, result=result)

    def detect_dict(
        self, d: t.Dict[JSONType, JSONType], *, path: Path, result: Result
    ) -> TypeInfo:
        props = {}
        sigs = []

        for k, v in d.items():
            path.append(k)
            subinfo = props[k] = self.detect(v, path=path, result=result)
            sigs.append((k, subinfo.sig))
            path.pop()

        sig = frozenset(sigs)
        uid = result._uid_map[sig]
        info = result.get(uid, default=SENTINEL)

        if info is not SENTINEL:
            info._others.append((path[:], d))
            return result.registry[uid]

        return result.add(uid, Object(sig, path=path[:], raw=d, props=props))

    def detect_list(
        self, d: t.List[JSONType], *, path: Path, result: Result
    ) -> TypeInfo:
        # TODO: zero
        # TODO: optional/union
        members = set()
        for i, x in enumerate(d):
            path.append(str(i))
            members.add(self.detect(x, path=path, result=result))
            path.pop()

        if len(members) == 0:
            item = ZERO
        elif len(members) == 1:
            item = members.pop()
        else:
            item = Union(members)  # xxx:

        base = list
        sig = tuple([base, item.sig])
        uid = result._uid_map[sig]
        return result.add(
            uid, Container(sig, base=base, item=item, path=path[:], raw=d)
        )

    def detect_primitive(self, d: t.Any, *, path: Path, result: Result) -> TypeInfo:
        sig = type(d)
        uid = result._uid_map[sig]
        if uid in result:
            return result.registry[uid]

        return result.add(uid, Primitive(sig, path=path[:], raw=d, type_=type(d)))


def detect(d: JSONType, *, detector=Detector()):
    path: Path = []
    result = Result()
    detector.detect(d, path=path, result=result)
    return result


def show(d):
    print("-")
    r = detect(d)
    print(r)
    for info in r.history:
        print(" ", info)
