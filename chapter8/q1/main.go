package main

import "fmt"

func main() {
	x := new(int)
	fmt.Println(&x)
}

// Result
// 0xc42000c028
