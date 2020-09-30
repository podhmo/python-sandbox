package main

import (
	"encoding/json"
	"m/shape"
	"os"
	"reflect"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// type Kind uint

// const (
// 	Invalid Kind = iota
// 	Bool
// 	Int
// 	Int8
// 	Int16
// 	Int32
// 	Int64
// 	Uint
// 	Uint8
// 	Uint16
// 	Uint32
// 	Uint64
// 	Uintptr
// 	Float32
// 	Float64
// 	Complex64
// 	Complex128
// 	Array
// 	Chan
// 	Func
// 	Interface
// 	Map
// 	Ptr
// 	Slice
// 	String
// 	Struct
// 	UnsafePointer
// )

type Person struct {
	Title string `json:"title"` // required
	Age   int    `json:"age"`
}

func ListPerson() []Person {
	return nil
}
func GetPerson() []Person {
	return nil
}

func Emit(ob interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(ob)
}

func FieldName(s shape.Struct, i int) string {
	name := s.Fields.Keys[i]
	if v, ok := s.Tags[i].Lookup("json"); ok {
		name = strings.SplitN(v, ",", 2)[0] // todo: omitempty, inline
	}
	return name
}

func Transform(s shape.Shape) interface{} {
	switch s := s.(type) {
	case shape.Struct:
		schema := openapi3.NewObjectSchema()
		for i, v := range s.Fields.Values {
			// todo: json tag
			name := FieldName(s, i)
			switch v.GetKind() {
			case reflect.String:
				f := openapi3.NewStringSchema()
				schema.Properties[name] = &openapi3.SchemaRef{Value: f}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				f := openapi3.NewIntegerSchema()
				schema.Properties[name] = &openapi3.SchemaRef{Value: f}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64: // Uintptr
				f := openapi3.NewIntegerSchema()
				schema.Properties[name] = &openapi3.SchemaRef{Value: f}
			default:
				panic(v)
			}
		}
		return schema
	default:
		panic(s)
	}
}

func main() {
	Emit(Transform(shape.Extract(Person{})))
}
