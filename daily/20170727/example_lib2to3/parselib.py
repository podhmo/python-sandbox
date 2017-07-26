import sys
from lib2to3 import pytree
from lib2to3 import pygram
from lib2to3.pgen2 import driver
from lib2to3.pgen2 import token

default_driver = driver.Driver(pygram.python_grammar_no_print_statement, convert=pytree.convert)


def parse_from_string(code, parser_driver=default_driver, *, debug=True):
    return parser_driver.parse_string(code, debug=debug)


def parse_from_file(filename, parser_driver=default_driver, *, debug=True):
    return parser_driver.parse_file(filename, debug=debug)


def dump_tree(tree, target_stream=sys.stdout):
    dumper = PyTreeDumper(target_stream)
    dumper.visit(tree)


def node_name(node):
    # Nodes with values < 256 are tokens. Values >= 256 are grammar symbols.
    if node.type < 256:
        return token.tok_name[node.type]
    else:
        return pygram.python_grammar.number2symbol[node.type]


def node_fullname(node):
    typ = pytree.type_repr(node.type)
    if typ == "funcdef":
        return "{}[name={}]".format(node_name(node), repr(node.children[1].value))
    elif typ == "classdef":
        return "{}[name={}]".format(node_name(node), repr(node.children[1].value))
    elif typ == "typedargslist":
        return "{}[args={}]".format(node_name(node), " ".join([repr(c.value) for c in node.children]))
    else:
        return node_name(node)


def dump_node_to_string(node):
    if not hasattr(node, "children"):  # leaf
        fmt = '{name}({value}) [lineno={lineno}, column={column}, prefix={prefix}]'
        return fmt.format(
            name=node_fullname(node),
            value=repr(node),
            lineno=node.lineno,
            column=node.column,
            prefix=repr(node.prefix)
        )
    else:
        fmt = '{node} [{len} children]'
        return fmt.format(
            node=node_fullname(node),
            len=len(node.children),
        )


# from yapf
class PyTreeVisitor(object):
    def visit(self, node):
        method = 'visit_{0}'.format(node_name(node))
        if hasattr(self, method):
            # Found a specific visitor for this node
            getattr(self, method)(node)
        else:
            if isinstance(node, pytree.Leaf):
                self.default_leaf_visit(node)
            else:
                self.default_node_visit(node)

    def default_node_visit(self, node):
        for child in node.children:
            self.visit(child)

    def default_leaf_visit(self, leaf):
        pass


class PyTreeDumper(PyTreeVisitor):
    def __init__(self, target_stream=sys.stdout):
        self._target_stream = target_stream
        self._current_indent = 0

    def _dump_string(self, s):
        self._target_stream.write('{0}{1}\n'.format(' ' * self._current_indent, s))

    def default_node_visit(self, node):
        self._dump_string(dump_node_to_string(node))
        self._current_indent += 2
        super(PyTreeDumper, self).default_node_visit(node)
        self._current_indent -= 2

    def DefaultLeafVisit(self, leaf):
        self._dump_string(dump_node_to_string(leaf))


class StrictVisitor(PyTreeVisitor):
    def default_node_visit(self, node):
        for child in node.children:
            self.visit(child)
        method = 'visit_{0}'.format(node_name(node))
        if not hasattr(self, method):
            raise NotImplemented(method)

    def default_leaf_visit(self, node):
        method = 'visit_{0}'.format(node_name(node))
        if not hasattr(self, method):
            raise NotImplemented(method)
