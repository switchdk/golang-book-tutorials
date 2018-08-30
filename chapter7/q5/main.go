package main

import "fmt"

func fib(n uint) uint {
	if n == 0 {
		return n
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	fibN := 20
	for i := 0; i < fibN; i++ {
		fmt.Println(fib(uint(i)))
	}
}

// Result
// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// 55
// 89
// 144
// 233
// 377
// 610
// 987
// 1597
// 2584
// 4181
