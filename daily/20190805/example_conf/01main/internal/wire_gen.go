// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package internal

import (
	"m/bar"
	"m/boo"
	"m/conf"
	"m/foo"
)

// Injectors from wire.go:

func InitializeRoot(c *conf.Config) *Root {
	fooFoo := foo.New(c)
	barBar := bar.New(c)
	booBoo := boo.New(c)
	root := NewRoot(fooFoo, barBar, booBoo)
	return root
}
