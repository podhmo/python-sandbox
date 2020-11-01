# -*- coding:utf-8 -*-
# this is auto-generated by swagger-marshmallow-codegen
from marshmallow import (
    Schema,
    fields,
)
from swagger_marshmallow_codegen.schema import AdditionalPropertiesSchema


class S(Schema):
    x = fields.Nested(lambda: SX)
    y = fields.Nested(lambda: SY)
    z = fields.Nested(lambda: D)


class SY(AdditionalPropertiesSchema):
    value = fields.Integer()

    class Meta:
        additional_field = fields.Nested(lambda: O)



class SX(AdditionalPropertiesSchema):
    value = fields.Integer()

    class Meta:
        additional_field = fields.Field()



class O(Schema):
    name = fields.String()


class D(AdditionalPropertiesSchema):
    value = fields.Integer()

    class Meta:
        additional_field = fields.Field()
