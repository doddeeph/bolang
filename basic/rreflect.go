package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	fmt.Println("type:", reflect.TypeOf(number))
	fmt.Println("type:", reflect.ValueOf(number).Type())
	fmt.Println("value:", reflect.ValueOf(number).Interface())
	fmt.Println("value:", reflect.ValueOf(number).Interface().(int))

	if reflect.ValueOf(number).Kind() == reflect.Int {
		fmt.Println("value:", reflect.ValueOf(number).Int())
	}

	var m = &murid{"wick", 2}
	m.getPropertyInfo()

	fmt.Println(m.Name)
	var reflectValue = reflect.ValueOf(m)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("john"),
	})
	fmt.Println(m.Name)
}

// catatan property-property tersebut bermodifier public
type murid struct {
	Name  string
	Grade int
}

func (m *murid) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(m)
	var reflectType = reflectValue.Type()

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectType = reflectValue.Type()
	}

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama\t:", reflectType.Field(i).Name)
		fmt.Println("tipe\t:", reflectType.Field(i).Type)
		fmt.Println("nilai\t:", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

// harus bermodifier public
func (m *murid) SetName(name string) {
	m.Name = name
}
