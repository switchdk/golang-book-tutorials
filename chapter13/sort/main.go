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

const unsortedList string = "unsortedList.txt"
const sortedList string = "sortedList.txt"

// ByValue List with integers to sort
type ByValue []int

func (bv ByValue) Len() int {
	return len(bv)
}

func (bv ByValue) Less(i, j int) bool {
	return bv[i] < bv[j]
}

func (bv ByValue) Swap(i, j int) {
	bv[i], bv[j] = bv[j], bv[i]
}

func sortList() {

	values, err := ioutil.ReadFile(unsortedList)
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

	file := openFile(sortedList)
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
	file := openFile(unsortedList)
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
