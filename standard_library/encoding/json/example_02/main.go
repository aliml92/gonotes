package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DataType struct {
	Int            int       `json:"int,omitempty"`
	String         string    `json:"string,omitempty"`
	Bool           bool      `json:"bool,omitempty"`  
	SliceOfInts    []int     `json:"slice_of_ints,omitmepty"`
	SliceOfBools   []bool    `json:"slice_of_bools,omitmepty"`
	SliceOfStrings []string  `json:"slice_of_strings,omitmepty"`
	ArrayOfInts    [2]int    `json:"array_of_ints,omitempty"`
	ArrayOfStrings [2]string `json:"array_of_strings,omitempty"`                    
}

type DataTypeWithPointer struct {
	Int            *int       `json:"int,omitempty"`
	String         *string    `json:"string,omitempty"`
	Bool           *bool      `json:"bool,omitempty"`  
	SliceOfInts    *[]int     `json:"slice_of_ints,omitmepty"`
	SliceOfBools   *[]bool    `json:"slice_of_bools,omitmepty"`
	SliceOfStrings *[]string  `json:"slice_of_strings,omitmepty"`
	ArrayOfInts    *[2]int    `json:"array_of_ints,omitempty"`
	ArrayOfStrings *[2]string `json:"array_of_strings,omitempty"`                    
}

var dataType1 = `
{
	"int": 0,
	"string": "",
	"bool": false,
	"slice_of_ints": [],
	"slice_of_bools": [],
	"slice_of_strings": [],
	"array_of_ints": [],
	"array_of_strings": []
}  
`

var dataType2 = `
{
	"int": 0,
	"string": "",
	"bool": false,
	"slice_of_ints": [0,0],
	"slice_of_bools": [false, false],
	"slice_of_strings": ["",""],
	"array_of_ints": [0,0],
	"array_of_strings": ["",""]
}  
`
var dataType3 = `{}`

var dataType4 = `
{
	"slice_of_ints": [0,0],
	"slice_of_bools": [false, false],
	"slice_of_strings": ["",""],
	"array_of_ints": [0,0],
	"array_of_strings": ["",""]
}  
`

var dataType5 = `
{
	"int": 0,
	"string": "",
	"bool": false
}  
`


func example1(){
	var v1 DataType
	r := bytes.NewReader([]byte(dataType1))
	err := json.NewDecoder(r).Decode(&v1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType1: %v\n", v1)

	var v2 DataTypeWithPointer
	r2 := bytes.NewReader([]byte(dataType1))
	err = json.NewDecoder(r2).Decode(&v2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType2: %v\n", v2)
}

func example2(){
	var v1 DataType
	r := bytes.NewReader([]byte(dataType2))
	err := json.NewDecoder(r).Decode(&v1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType1: %v\n", v1)

	var v2 DataTypeWithPointer
	r2 := bytes.NewReader([]byte(dataType2))
	err = json.NewDecoder(r2).Decode(&v2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType2: %v\n", v2)
}

func example3(){
	var v1 DataType
	r := bytes.NewReader([]byte(dataType3))
	err := json.NewDecoder(r).Decode(&v1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType1: %v\n", v1)

	var v2 DataTypeWithPointer
	r2 := bytes.NewReader([]byte(dataType3))
	err = json.NewDecoder(r2).Decode(&v2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType2: %v\n", v2)
}

func example4(){
	var v1 DataType
	r := bytes.NewReader([]byte(dataType4))
	err := json.NewDecoder(r).Decode(&v1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType1: %v\n", v1)

	var v2 DataTypeWithPointer
	r2 := bytes.NewReader([]byte(dataType4))
	err = json.NewDecoder(r2).Decode(&v2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType2: %v\n", v2)
}

func example5(){
	var v1 DataType
	r := bytes.NewReader([]byte(dataType5))
	err := json.NewDecoder(r).Decode(&v1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType1: %v\n", v1)

	var v2 DataTypeWithPointer
	r2 := bytes.NewReader([]byte(dataType5))
	err = json.NewDecoder(r2).Decode(&v2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dataType2: %v\n", v2)
}

func main() {
	fmt.Println("==== example1() ====")
	example1()

	fmt.Println("==== example2() ====")
	example2()

	fmt.Println("==== example3() ====")
	example3()

	fmt.Println("==== example4() ====")
	example4()

	fmt.Println("==== example5() ====")
	example5()
}