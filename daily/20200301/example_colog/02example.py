print("[ trace ] logging this to stdout")
print("[ \x1b[0;36mdebug\x1b[0m ] logging this to stdout")
print("[  \x1b[0;32minfo\x1b[0m ] logging this to stdout")
print("[  \x1b[0;33mwarn\x1b[0m ] with fields foo=bar")
print("\x1b[0;31m[ error ]\x1b[0m with fields foo=bar")
print("\x1b[0;37;41m[ alert ]\x1b[0m with fields foo=bar")