# this is auto-generated by swagger-marshmallow-codegen
from __future__ import annotations
from marshmallow import (
    Schema,
    fields,
    RAISE,
)
from swagger_marshmallow_codegen.validate import ItemsRange, Unique


class ExampleSchema(Schema):
    array = fields.List(
        fields.String(),
        required=True,
        description="Example array requiring at least one element",
        validate=[
            ItemsRange(min=1, max=None, min_inclusive=True, max_inclusive=True),
            Unique(),
        ],
    )

    class Meta:
        unknown = RAISE


schema = ExampleSchema()
schema.load({"array": ["string1", "string2"]})
