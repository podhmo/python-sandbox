# comment mesage

export TARGET ?=

default: hello bye

hello:
	echo hello
	echo hello

bye: hello
	echo $(shell echo bye)

