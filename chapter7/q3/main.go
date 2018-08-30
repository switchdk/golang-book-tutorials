package main

import "fmt"

func greatestNumber(intList ...int) int {
	if len(intList) == 0 {
		return 0
	} else if len(intList) == 1 {
		return intList[0]
	} else {
		gn := intList[0]
		for _, i := range intList {
			if i > gn {
				gn = i
			}
		}
		return gn
	}
}

func main() {
	fmt.Println(greatestNumber(1, 2, 42, 4, 5))
	// Result
	// 42

	fmt.Println(greatestNumber())
	// Result
	// 0

	fmt.Println(greatestNumber(5))
	// Result
	// 5
}
