package main

import "fmt"

func sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	result := sum(x)
	fmt.Println(result)
}

// Result
// 15
// Answer
// sum(numbers []int) int
