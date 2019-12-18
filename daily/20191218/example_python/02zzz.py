from __future__ import annotations
import typing as t
import typing_extensions as tx
from prestring import Module


class Emittable(tx.Protocol):
    def emit(self, *, m: Module) -> Module:
        ...


class Code:
    emit = None

    def __getattr__(self, name: str) -> Attr:
        return Attr(name, self)


class Attr:
    def __init__(self, name: str, co: t.Union[Emittable, Code, Attr]):
        self.name = name
        self._co = co
