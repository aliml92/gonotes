package main

import (
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


func set() {
	dt := DataType{
		Int:            0,
		String:         "",
		Bool:           false,
		SliceOfInts:    []int{},
		SliceOfBools:   []bool{},
		SliceOfStrings: []string{},
		ArrayOfInts:    [2]int{},
		ArrayOfStrings: [2]string{},
	}

	dtwp := DataTypeWithPointer{
		Int:            &dt.Int,
		String:         &dt.String,
		Bool:           &dt.Bool,
		SliceOfInts:    &dt.SliceOfInts,
		SliceOfBools:   &dt.SliceOfBools,
		SliceOfStrings: &dt.SliceOfStrings,
		ArrayOfInts:    &dt.ArrayOfInts,
		ArrayOfStrings: &dt.ArrayOfStrings,
	}

	data, err := json.Marshal(dt)
	if err != nil {
		fmt.Printf("Error marshalling DataType: %v\n", err)
		return
	}
	fmt.Printf("Marshalled DataType: %s\n", string(data))

	dataWithPointer, err := json.Marshal(dtwp)
	if err != nil {
		fmt.Printf("Error marshalling DataTypeWithPointer: %v\n", err)
		return
	}
	fmt.Printf("Marshaled DataTypeWithPointer: %v\n", string(dataWithPointer))
}

func unset() {
	dt := DataType{}

	dtwp := DataTypeWithPointer{}

	data, err := json.Marshal(dt)
	if err != nil {
		fmt.Printf("Error marshalling DataType: %v\n", err)
		return
	}
	fmt.Printf("Marshalled DataType: %s\n", string(data))

	dataWithPointer, err := json.Marshal(dtwp)
	if err != nil {
		fmt.Printf("Error marshalling DataTypeWithPointer: %v\n", err)
		return
	}
	fmt.Printf("Marshaled DataTypeWithPointer: %v\n", string(dataWithPointer))
}


func mixed() {
	dt := DataType{}

	dtwp := DataTypeWithPointer{
		Int:            &dt.Int,
		String:         &dt.String,
		Bool:           &dt.Bool,
		SliceOfInts:    &dt.SliceOfInts,
		SliceOfBools:   &dt.SliceOfBools,
		SliceOfStrings: &dt.SliceOfStrings,
		ArrayOfInts:    &dt.ArrayOfInts,
		ArrayOfStrings: &dt.ArrayOfStrings,
	}

	data, err := json.Marshal(dt)
	if err != nil {
		fmt.Printf("Error marshalling DataType: %v\n", err)
		return
	}
	fmt.Printf("Marshalled DataType: %s\n", string(data))

	dataWithPointer, err := json.Marshal(dtwp)
	if err != nil {
		fmt.Printf("Error marshalling DataTypeWithPointer: %v\n", err)
		return
	}
	fmt.Printf("Marshaled DataTypeWithPointer: %v\n", string(dataWithPointer))
}

func main() {
	fmt.Println("===== set() ====")
	set()
	fmt.Println("===== unset() ====")
	unset()
	fmt.Println("===== mixed() ====")
	mixed()		
}