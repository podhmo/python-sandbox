# -*- coding:utf-8 -*-
# this is auto-generated by swagger-marshmallow-codegen
from marshmallow import (
    Schema,
    fields
)
from swagger_marshmallow_codegen.schema import AdditionalPropertiesSchema


class Person(Schema):
    name = fields.String()


class Data(AdditionalPropertiesSchema):

    class Meta:
        additional_field = fields.List(fields.Nested('Person'))
