import typing as t
import typing_extensions as tx
import mypy_extensions as mx


class HasA(tx.Protocol):
    # a: str
    @property
    def a(self):
        pass

    @a.setter
    def a(self, v):
        pass


class HasB(tx.Protocol):
    b: str


T = t.TypeVar("T")


class AB(HasA, HasB):
    pass


def use(ob: AB) -> None:
    print(ob.a, ob.b)


Params = mx.TypedDict("Params", {"a": str, "b": str})


def run(d: Params) -> None:
    class Ob(AB):
        z = "boo"
        # a = d["a"]
        b = d["b"]

    ob = Ob()
    ob.a = d["a"]
    use(ob)
    print(ob.z)


def main() -> None:
    run({"a": "foo", "b": "bar"})


if __name__ == "__main__":
    main()
