from __future__ import annotations
import typing as t
import typing_extensions as tx
from prestring.text import Module
from prestring.go import goname
from egoist.go.resolver import get_resolver
from walker import walk
from metashape.declarative import field, MISSING


class Person:
    name: str
    age: int
    info: Info
    Father: Person = field(metadata={"pointer": True})
    Mother: Person = field(metadata={"pointer": True})


class Person2:
    name: str
    age: int
    info: Info = field(metadata={"inline": True})


class Info:
    memo: str


def metadata(
    *, inline: bool = False, pointer: bool = False, required: bool = True
) -> t.Dict[str, t.Any]:
    # todo: required false?
    d: Metadata = {"inline": inline, "pointer": pointer, "required": required}
    return d  # type: ignore


class Metadata(tx.TypedDict, total=False):
    inline: bool
    pointer: bool
    required: bool
    default: t.Any


def run() -> Module:
    r = get_resolver()
    m = Module(indent="\t")
    classes = [Person, Person2]

    for item in walk(classes):
        m.stmt(f"type {goname(item.cls.__name__)} struct {{")
        with m.scope():
            for name, typeinfo, _metadata in item.fields:
                metadata = t.cast(Metadata, _metadata)
                if metadata.get("default") == MISSING:
                    metadata.pop("default")

                try:
                    gotype = r.resolve_gotype(typeinfo.normalized)  # todo: pointer
                except KeyError:
                    gotype = goname(typeinfo.normalized.__name__)

                if metadata.get("pointer", False):
                    gotype = f"*{gotype}"

                if metadata.get("inline", False):
                    m.append(gotype)
                else:
                    m.append(f"{goname(name)} {gotype}")

                if metadata:
                    m.stmt(f"  // {metadata}")
                else:
                    m.stmt("")

        m.stmt("}")
        m.sep()
    return m


def gofmt(code: str) -> str:
    import subprocess
    import tempfile

    with tempfile.TemporaryFile("w+") as wf:
        wf.write(code)
        wf.seek(0)
        p = subprocess.run(
            ["gofmt"], stdin=wf, stdout=subprocess.PIPE, text=True, check=True
        )
        return p.stdout


print(gofmt(str(run())))
