package main

import (
  "fmt"
  // "os"
  "io/ioutil"
)

func main() {
  // file, err := os.Open("test.txt")
  // if err != nil {
  //   fmt.Println("Fool")
  //   return
  // }
  // defer file.Close()
  //
  // stat, err := file.Stat()
  // if err != nil {
  //   return
  // }
  //
  // fmt.Println("stat", stat.Size())
  //
  // bs := make([]byte, stat.Size())
  // _, err = file.Read(bs)
  // if err != nil {
  //   return
  // }
  // fmt.Println(string(bs))

  // Shortcut version
  bs, err := ioutil.ReadFile("test.txt")
  if err != nil {
    fmt.Println("fool")
    return
  }
  fmt.Println(string(bs))
}
