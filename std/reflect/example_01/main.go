package main

import (
	"fmt"
	"reflect"
)

func prettyPrint(i interface{}) {
	val := reflect.ValueOf(i)
	typ := reflect.TypeOf(i)

	fmt.Printf("Type: %s\n", typ)
	fmt.Printf("Value: %v\n", val)

	if val.Kind() == reflect.Struct {
		fmt.Println("Fields:")
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			value := val.Field(i)
			fmt.Printf("  %s: %v\n", field.Name, value.Interface())
		}
	}
}

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "John", Age: 30}
	prettyPrint(p)

	x := 42
	prettyPrint(x)
}
