package main

import "fmt"

func main() {
	arr := [5]string{"Hello", "my", "name", "is", "Gopher"}

	fmt.Println(arr[3])
	// Result
	// is

	for _, token := range arr {
		fmt.Println(token)
	}
	// Result
	// Hello
	// my
	// name
	// is
	// Gopher
}
