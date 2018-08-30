package main

import "fmt"

func swap(x, y *int) {
	tmp := new(int)
	*tmp = *x
	*x = *y
	*y = *tmp
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x:", x, "y:", y)
}

// Result
// x: 2 y: 1
