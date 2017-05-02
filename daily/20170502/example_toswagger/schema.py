# -*- coding:utf-8 -*-
# this is auto-generated by swagger-marshmallow-codegen
from marshmallow import (
    Schema,
    fields
)


class ArrayItem(Schema):
    boolean = fields.Boolean()
    integer = fields.Integer()
    null_value = fields.Field(allow_none=True)


class Object(Schema):
    array = fields.List(fields.Nested('ArrayItem'), required=True)
    key = fields.String(required=True)


class Top(Schema):
    content = fields.String(required=True)
    json = fields.List(fields.String(), required=True)
    object = fields.Nested('Object', required=True)
    paragraph = fields.String(required=True)
    yaml = fields.List(fields.String(), required=True)
