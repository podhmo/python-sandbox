import string
import pygtrie
import numpy.random

cs = numpy.array(list(str(string.printable)), dtype="|S1")
size = 50
N = 100000


def gen():
    words = numpy.random.choice(cs, [N, size])
    t = pygtrie.StringTrie()
    for word in words:
        t[str(word)] = str(word)
    return t


t = gen()
import psutil
print(psutil.Process().memory_info())
