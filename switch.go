package main

import (
  "fmt"
  "time"
)

func main() {
  i := 2
  fmt.Println("Write", i, "as:")
  switch i {
  case 1:
    fmt.Println("One")
  case 2:
    fmt.Println("Two")
  case 3:
    fmt.Println("Three")
  case 4:
    fmt.Println("Four")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  default:
    fmt.Println("it's a weekday")
  }

  t := time.Now()

  switch {
    case t.Hour() < 12:
      fmt.Println("It's before noon")
    default:
      fmt.Println("It's after noon")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm an int")
    default:
      fmt.Println("Don't know type %T\n", t)
    }
  }

  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")
}
