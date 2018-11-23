package main

import "fmt"

func incrementGene() func() int{
  x := 0
  return func () int{
    x++
    return x
  }
}

func circleArea(pi float64) func(radius float64) float64{
  return func(radius float64) float64{
    return pi * radius * radius
  }
}

func main(){
  count := incrementGene()
  fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	fmt.Println(c1(3))

}


