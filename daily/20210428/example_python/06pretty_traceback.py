import pretty_traceback

pretty_traceback.install()


def f0():
    f1()


def f1():
    f2()


def f2():
    f3()


def f3():
    f4()


def f4():
    f5()


def f5():
    f6()


def f6():
    x = locals()
    f7()


def run():
    f0()


run()
