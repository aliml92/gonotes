package main

import "fmt"

func main() {
	str := "Hello, 世界"

	// len over string value gives the length of slice of bytes 
	// not length of characters
	fmt.Println("str length: ", len(str))
	
	// indexing of string value returns byte value, instead of character
	// type will be byte (uint8)
	for i := 0; i < len(str); i++ {
		fmt.Printf("type: %T ", str[i])
		fmt.Printf("value: %v ", str[i])
		fmt.Printf("string value: %s\n", string(str[i]))
	}

	// ranging over string value spits out code points aka runes
	//type will be rune (int32)
	for index, runeValue := range str {
		fmt.Printf("type: %T ", runeValue)
		fmt.Printf("value: %v ", runeValue)
		fmt.Printf("at pos: %d\n", index)
	}
}