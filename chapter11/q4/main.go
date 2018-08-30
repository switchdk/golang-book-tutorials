package main

import "fmt"
import m "golang-book/chapter11/q4/math"

func main() {
	xs := []float64{1, 7, 3, 1, 5, 6, 7, 8}

	avg := m.Average(xs)
	fmt.Println("Average:", avg)

	min := m.Min(xs)
	fmt.Println("Minimum:", min)

	max := m.Max(xs)
	fmt.Println("Maximum:", max)
}
