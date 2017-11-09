import gc
import math
import psutil
import pandas as pd
import numpy as np
import objgraph


def convert_size(size_bytes):
    if size_bytes == 0:
        return "0B"
    size_name = ("B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB")
    i = int(math.floor(math.log(size_bytes, 1024)))
    p = math.pow(1024, i)
    s = round(size_bytes / p, 2)
    return "%s %s" % (s, size_name[i])


p = psutil.Process()
print("========================================")
print("before allocated", convert_size(p.memory_info().rss))
objgraph.show_growth()
df = pd.DataFrame({"foo": np.arange(10000000)})
print("----------------------------------------")
df.info()
print("----------------------------------------")
print("========================================")
print("after allocated", convert_size(p.memory_info().rss))
objgraph.show_growth()
del df
print("========================================")
print("after delete variable", convert_size(p.memory_info().rss))
objgraph.show_growth()
gc.collect()
print("========================================")
print("after gc", convert_size(p.memory_info().rss))
objgraph.show_growth()
