from collections import ChainMap, defaultdict
import logging
logger = logging.getLogger(__name__)


class NameConflict(Exception):
    pass


class OrphanResourceIsFound(Exception):
    pass


class Resource:
    def __init__(self, coll, q=None):
        self.result = None
        self.mappings = {}

        self.coll = coll
        self.q = q or {}

    def in_(self, name, ids):
        if name in self.q and "$in" in self.q[name]:
            return self
        q = self.q.copy()
        q[name] = {"$in": ids}
        return self.__class__(self.coll, q=q)

    def __iter__(self):
        if self.result is not None:
            return iter(self.result)

        logger.debug("fetch %s %s", self.coll, self.q)
        self.result = []

        def iterate(append):
            for row in iter(self.coll.find(self.q)):
                append(row)
                yield row

        return iterate(self.result.append)

    def bind_one(self, name, subresource, rel):
        return bind_one(self, name, subresource, rel)

    def bind_many(self, name, subresource, rel):
        return bind_many(self, name, subresource, rel)


def bind_one(resource, name, subresource, rel):
    key, subkey = [k.split(".", 1)[-1] for k in rel.split("==")]  # xxx
    return BoundOneResource(resource, name, subresource, key=key, subkey=subkey)


def bind_many(resource, name, subresource, rel):
    key, subkey = [k.split(".", 1)[-1] for k in rel.split("==")]  # xxx
    return BoundManyResource(resource, name, subresource, key=key, subkey=subkey)


class _BoundResource:
    def __init__(self, resource, name, subresource, key, subkey):
        self.resource = resource
        self.name = name
        self.subkey = subkey
        self.key = key
        self.subresource = subresource

    def bind_one(self, name, subresource, rel):
        return bind_one(self, name, subresource, rel)

    def bind_many(self, name, subresource, rel):
        return bind_many(self, name, subresource, rel)


class BoundOneResource(_BoundResource):
    def __iter__(self):
        gk = self.key
        sk = self.subkey
        name = self.name
        mapping = {sr[sk]: sr for sr in self.subresource}
        for r in self.resource:
            yield ChainMap({name: mapping[r[gk]]}, r)


class BoundManyResource(_BoundResource):
    def __iter__(self):
        gk = self.key
        sk = self.subkey
        name = self.name
        mapping = defaultdict(list)
        for sr in self.subresource:
            mapping[sr[sk]].append(sr)
        for r in self.resource:
            yield ChainMap({name: mapping[r[gk]]}, r)
