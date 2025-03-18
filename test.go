package main

import (
  "fmt"
  "math"
) 

type Circle struct {
  x, y, r float64
}

type Shape interface {
  area() float64
}

// Area is a method on the Shape interface. It will be implemented by Circle using the receiver type.
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}

func main() {
  c := Circle{x: 0, y: 0, r: 5}
  // Pass a pointer of the Circle struct to totalArea
  fmt.Println(totalArea(&c))
}

