# this is auto-generated by swagger-marshmallow-codegen
from __future__ import annotations
from marshmallow import (
    Schema,
    fields,
)


class Person(Schema):
    name = fields.String()
    age = fields.Integer()
    father = fields.Nested(lambda: Person())
    mother = fields.Nested(lambda: Person())
