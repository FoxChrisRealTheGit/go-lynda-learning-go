package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}

	fmt.Println(colors)

	colors = append(colors, "Purple")

	fmt.Println(colors)

	//removes first
	// colors = append(colors[1:len(colors)] )
	// colors = append(colors[1:] )

	//removes last
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 52
	numbers[1] = 2

	numbers[2] = 45

	numbers[3] = 4

	numbers[4] = 45
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)

	fmt.Println(numbers)

}
