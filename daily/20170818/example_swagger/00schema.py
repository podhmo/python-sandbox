# -*- coding:utf-8 -*-
# this is auto-generated by swagger-marshmallow-codegen
from marshmallow import (
    Schema,
    fields
)
from swagger_marshmallow_codegen.schema import PrimitiveValueSchema


class S(Schema):
    person = fields.Nested('Person')
    nperson = fields.Field('Nperson')


class Person(Schema):
    name = fields.String()
    age = fields.Integer()


class Nperson(Person):
    pass


class Nperson2(PrimitiveValueSchema):
    class schema_class(Schema):
        value = fields.Field(allow_none=True)
