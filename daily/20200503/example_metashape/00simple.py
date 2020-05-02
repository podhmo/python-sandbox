from emit import emit
from prestring.go.gofmt import gofmt


class Person:
    """this is person object"""
    name: str
    age: int


print(gofmt(emit([Person]), always=False))
