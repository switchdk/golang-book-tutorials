package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type ByValue []int

func (this ByValue) Len() int {
	return len(this)
}

func (this ByValue) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this ByValue) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func sortList() {

	values, err := ioutil.ReadFile("to_sort1.txt")
	if err != nil {
		fmt.Println("Fool")
		return
	}

	number := make([]byte, 0)
	var intList []int

	for _, v := range values {
		if v != 10 {
			number = append(number, v)
		} else {
			x, err := strconv.Atoi(string(number))
			if err != nil {
				fmt.Println("String to Int conversion failed", err)
			}
			intList = append(intList, x)
			number = nil
		}
	}

	sort.Sort(ByValue(intList))

	file := openFile("sorted.txt")
	defer file.Close()

	for _, i := range intList {
		_, err := file.WriteString(fmt.Sprintln(i))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func createList(size int) {
	file := openFile("to_sort1.txt")
	defer file.Close()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		_, err := file.WriteString(fmt.Sprintln(strconv.Itoa(rand.Intn(size))))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func openFile(filename string) os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	return *file
}

func main() {
	maxSize := flag.Int("size", 10, "the maximum size of number")
	flag.Parse()
	createList(*maxSize)
	sortList()
}
