package main

import "fmt"

func main() {
  ch := make(chan int, 2)

  ch <- 0
  ch <- 1
  // ch <- 2

  a := <- ch
  b := <- ch

  fmt.Println(a, b)
}
