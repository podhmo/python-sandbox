import json
from collections import namedtuple
cmd = namedtuple("cmd", "op, v, v1")

# todo: move


def unpatch(src, dst, *, verbose=False):
    r = Walker().walk(src, dst)
    rows = merge(r)

    if not verbose:
        for row in rows:
            row.pop("from", None)
            if row["op"] == "remove":
                row.pop("value", None)
            yield row
    else:
        yield from rows


def merge(r):
    if r is None:
        return []
    if not hasattr(r, "keys"):
        yield {"path": "", "op": r.op, "value": r.v, "from": r.v1}
    else:
        for k, v in r.items():
            for sv in merge(v):
                yield {
                    "path": f"/{k}/{sv['path'].lstrip('/')}".rstrip("/"),
                    "op": sv["op"],
                    "value": sv["value"],
                    "from": sv.pop("from", None),
                }


class Walker:
    def __init__(self):
        self.move_map = {}  # todo:

    def walk(self, src, dst):
        # xxx: src and dst is None
        if hasattr(src, "keys"):
            return self._walk_dict(src, dst)
        elif isinstance(src, (list, tuple)):
            return self._walk_list(src, dst)
        else:
            return self._walk_atom(src, dst)

    def _walk_dict(self, src, dst):
        r = {}
        for k, v in src.items():
            if k in dst:
                r[k] = self.walk(v, dst[k])
            else:
                r[k] = cmd("remove", v=v, v1=None)
        for k, v in dst.items():
            if k in r:
                continue
            r[k] = cmd("add", v=v, v1=None)
        return r

    def _walk_list(self, src, dst):
        r = {}
        n = min(len(src), len(dst))
        for i in range(n):
            r[str(i)] = self.walk(src[i], dst[i])

        if n == len(dst):
            for i in range(n, len(src)):
                r[str(i)] = cmd(op="remove", v=src[i])
        else:
            for i in range(n, len(dst)):
                r[str(i)] = cmd(op="add", v=dst[i])
        return r

    def _walk_atom(self, src, dst):
        if src is None:
            return cmd(op="add", v=dst)
        elif dst is None:
            return cmd(op="remove", v=src)
        elif src != dst:
            return cmd(op="replace", v=dst, v1=src)
        else:
            return None


import unittest  # noqa


class Tests(unittest.TestCase):
    def _callFUT(self, src, dst):
        return unpatch(src, dst)

    def test00(self):
        src = {"name": "foo"}
        dst = {"name": "foo"}

        got = self._callFUT(src, dst)
        want = []
        self.assertListEqual(
            sorted([json.dumps(x, sort_keys=True) for x in got]),
            sorted([json.dumps(x, sort_keys=True) for x in want])
        )

    def test01(self):
        src = {"name": "foo"}
        dst = {"name": "bar"}

        got = self._callFUT(src, dst)
        want = [{"op": "replace", "path": "/name", "value": "bar"}]
        self.assertListEqual(
            sorted([json.dumps(x, sort_keys=True) for x in got]),
            sorted([json.dumps(x, sort_keys=True) for x in want])
        )

    def test02(self):
        src = {"name": "foo"}
        dst = {}

        got = self._callFUT(src, dst)
        want = [{"op": "remove", "path": "/name"}]
        self.assertListEqual(
            sorted([json.dumps(x, sort_keys=True) for x in got]),
            sorted([json.dumps(x, sort_keys=True) for x in want])
        )

    def test03(self):
        src = {}
        dst = {"name": "bar"}

        got = self._callFUT(src, dst)
        want = [{"op": "add", "path": "/name", "value": "bar"}]
        self.assertListEqual(
            sorted([json.dumps(x, sort_keys=True) for x in got]),
            sorted([json.dumps(x, sort_keys=True) for x in want])
        )

    def test04(self):
        src = {"point0": {"value": 10}}
        dst = {"point1": {"value": 10}}

        got = self._callFUT(src, dst)
        # todo: move
        want = [
            {
                "op": "remove",
                "path": "/point0",
            },
            {
                "op": "add",
                "path": "/point1",
                "value": {
                    "value": 10
                },
            },
        ]
        self.assertListEqual(
            sorted([json.dumps(x, sort_keys=True) for x in got]),
            sorted([json.dumps(x, sort_keys=True) for x in want])
        )

    def test05(self):
        src = {"point0": {"value": 10}}
        dst = {"point0": {"value": 20}}

        got = self._callFUT(src, dst)
        want = [
            {
                "op": "replace",
                "path": "/point0/value",
                "value": 20,
            },
        ]
        self.assertListEqual(
            sorted([json.dumps(x, sort_keys=True) for x in got]),
            sorted([json.dumps(x, sort_keys=True) for x in want])
        )

    def test06(self):
        src = {"person": {"name": "foo", "age": 20, "type": "P"}}
        dst = {"person": {"name": "bar", "nickname": "B", "type": "P"}}
        got = self._callFUT(src, dst)
        want = [
            {
                "op": "replace",
                "path": "/person/name",
                "value": "bar",
            },
            {
                "op": "remove",
                "path": "/person/age",
            },
            {
                "op": "add",
                "path": "/person/nickname",
                "value": "B",
            },
        ]
        self.assertListEqual(
            sorted([json.dumps(x, sort_keys=True) for x in got]),
            sorted([json.dumps(x, sort_keys=True) for x in want])
        )

    def test07(self):
        src = [{}, {"person": {"name": "foo", "age": 20, "type": "P"}}]
        dst = [{"person": {"name": "bar", "nickname": "B", "type": "P"}}, {}]
        got = self._callFUT(src, dst)
        want = [
            {
                "path": "/0/person",
                "op": "add",
                "value": {
                    "name": "bar",
                    "nickname": "B",
                    "type": "P"
                }
            },
            {
                "path": "/1/person",
                "op": "remove",
            },
        ]
        self.assertListEqual(
            sorted([json.dumps(x, sort_keys=True) for x in got]),
            sorted([json.dumps(x, sort_keys=True) for x in want])
        )


if __name__ == "__main__":
    unittest.main()
