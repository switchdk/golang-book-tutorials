package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("fool")
		return
	}
	fmt.Println(string(bs))
}
