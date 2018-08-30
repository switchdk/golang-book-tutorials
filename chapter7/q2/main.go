package main

import "fmt"

func half(i int) (int, bool) {
	return i / 2, i%2 == 0
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	for _, i := range x {
		result, even := half(i)
		fmt.Println("half(", i, ")", ", even:", result, even)
	}
}

// Result
// half( 1 ) , even: 0 false
// half( 2 ) , even: 1 true
// half( 3 ) , even: 1 false
// half( 4 ) , even: 2 true
// half( 5 ) , even: 2 false
