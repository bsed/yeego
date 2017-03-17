/**
 * Created by angelina-zf on 17/2/25.
 */
package yeeReflect_test

import (
	"testing"
	"yeego/yeeReflect"
	"reflect"
	"yeego"
)

type a struct {
}

func TestGetTypeFullName(t *testing.T) {
	yeego.Equal(yeeReflect.GetTypeFullName(reflect.TypeOf("")), "string")
	yeego.Equal(yeeReflect.GetTypeFullName(reflect.TypeOf(1)), "int")
	yeego.Equal(yeeReflect.GetTypeFullName(reflect.TypeOf(&a{})), "yeego/yeeReflect_test.a")
}

func TestIndirectType(t *testing.T) {
	yeego.Equal(yeeReflect.IndirectType(reflect.TypeOf(&a{})), reflect.TypeOf(a{}))
}

type GetAllFieldT1 struct {
	GetAllFieldT2
	A int
}

type GetAllFieldT2 struct {
	B string
}

func TestStructGetAllField(t *testing.T) {
	fileds := yeeReflect.StructGetAllField(reflect.TypeOf(&GetAllFieldT1{}))
	yeego.Equal(len(fileds), 3)
	yeego.Equal(fileds[0].Name, "GetAllFieldT2")
	yeego.Equal(fileds[0].Index, []int{0})
	yeego.Equal(fileds[1].Name, "A")
	yeego.Equal(fileds[1].Index, []int{1})
	yeego.Equal(fileds[2].Name, "B")
	yeego.Equal(fileds[2].Index, []int{0, 0})
}

type GetAllFieldT3 struct {
	A int
	B string
	C string
}

func TestGetTypeFullName2(t *testing.T) {
	fieldT3 := &GetAllFieldT3{
		A: 1,
		B: "string",
		C: "test",
	}
	data := make([]interface{}, 3)
	v := reflect.ValueOf(*fieldT3)
	fields := yeeReflect.StructGetAllField(reflect.TypeOf(fieldT3))
	for i := 0; i < len(fields); i++ {
		value := v.FieldByName(fields[i].Name)
		data[i] = value.Interface()
	}
	yeego.Equal(data, []interface{}{1, "string", "test"})
}
