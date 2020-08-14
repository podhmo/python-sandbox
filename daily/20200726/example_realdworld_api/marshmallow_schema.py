# -*- coding:utf-8 -*-
# this is auto-generated by swagger-marshmallow-codegen
from marshmallow import (
    Schema,
    fields,
)


class LoginUser(Schema):
    email = fields.String(required=True)
    password = fields.String(required=True)


class LoginUserRequest(Schema):
    user = fields.Nested('LoginUser', required=True)


class NewUser(Schema):
    username = fields.String(required=True)
    email = fields.String(required=True)
    password = fields.String(required=True)


class NewUserRequest(Schema):
    user = fields.Nested('NewUser', required=True)


class User(Schema):
    email = fields.String(required=True)
    token = fields.String(required=True)
    username = fields.String(required=True)
    bio = fields.String(required=True)
    image = fields.String(required=True)


class UserResponse(Schema):
    user = fields.Nested('User', required=True)


class UpdateUser(Schema):
    email = fields.String()
    token = fields.String()
    username = fields.String()
    bio = fields.String()
    image = fields.String()


class UpdateUserRequest(Schema):
    user = fields.Nested('UpdateUser', required=True)


class ProfileResponse(Schema):
    profile = fields.Nested('Profile', required=True)


class Profile(Schema):
    username = fields.String(required=True)
    bio = fields.String(required=True)
    image = fields.String(required=True)
    following = fields.Boolean(required=True)


class Article(Schema):
    slug = fields.String(required=True)
    title = fields.String(required=True)
    description = fields.String(required=True)
    body = fields.String(required=True)
    tagList = fields.List(fields.String(), required=True)
    createdAt = fields.DateTime(required=True)
    updatedAt = fields.DateTime(required=True)
    favorited = fields.Boolean(required=True)
    favoritesCount = fields.Integer(required=True)
    author = fields.Nested('Profile', required=True)


class SingleArticleResponse(Schema):
    article = fields.Nested('Article', required=True)


class MultipleArticlesResponse(Schema):
    articles = fields.List(fields.Nested('Article'), required=True)
    articlesCount = fields.Integer(required=True)


class NewArticle(Schema):
    title = fields.String(required=True)
    description = fields.String(required=True)
    body = fields.String(required=True)
    tagList = fields.List(fields.String())


class NewArticleRequest(Schema):
    article = fields.Nested('NewArticle', required=True)


class UpdateArticle(Schema):
    title = fields.String()
    description = fields.String()
    body = fields.String()


class UpdateArticleRequest(Schema):
    article = fields.Nested('UpdateArticle', required=True)


class Comment(Schema):
    id = fields.Integer(required=True)
    createdAt = fields.DateTime(required=True)
    updatedAt = fields.DateTime(required=True)
    body = fields.String(required=True)
    author = fields.Nested('Profile', required=True)


class SingleCommentResponse(Schema):
    comment = fields.Nested('Comment', required=True)


class MultipleCommentsResponse(Schema):
    comments = fields.List(fields.Nested('Comment'), required=True)


class NewComment(Schema):
    body = fields.String(required=True)


class NewCommentRequest(Schema):
    comment = fields.Nested('NewComment', required=True)


class TagsResponse(Schema):
    tags = fields.List(fields.String(), required=True)


class GenericErrorModel(Schema):
    errors = fields.Nested('GenericErrorModelErrors', required=True)


class GenericErrorModelErrors(Schema):
    body = fields.List(fields.String(), required=True)