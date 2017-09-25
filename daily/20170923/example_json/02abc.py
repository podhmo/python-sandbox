import datetime as dt
import extjson
import abc


class Stringer(abc.ABC):
    pass


Stringer.register(dt.datetime)
Stringer.register(dt.date)


@extjson.register(Stringer)
def encode_stringer(o):
    return str(o)


d = {
    "name": "foo",
    "birth": dt.datetime.now(),
    "birthday": dt.date.today(),
}
print(extjson.dumps(d))
