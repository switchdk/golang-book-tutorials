package main

import "fmt"

func makeEvenOddGenerator(startValue int) func() uint {
	i := uint(startValue)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	fmt.Println("Even")
	nextEven := makeEvenOddGenerator(0)
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("Odd")
	nextOdd := makeEvenOddGenerator(1)
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}

// Result
// Even
// 0
// 2
// 4
// Odd
// 1
// 3
// 5
