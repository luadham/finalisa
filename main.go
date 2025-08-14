package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello World")
}

func Init() {
  fmt.Println(1/0)
}
