package main

import "fmt"

func main() {
	x := new(int)
	fmt.Println(*x)
	*x = 1
	fmt.Println(*x)
}

// Result
// 0
// 1
