package main

import (
	"fmt"
	"strings"
)

func main() {
	str1:= "implicitly typed string"
	fmt.Printf("str1: %v:%T", str1, str1)
	var str2 = "explicitly typed string"
	fmt.Printf("str2: %v:%T", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	
}
