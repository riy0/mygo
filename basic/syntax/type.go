package main

import (
  "fmt"
  "strings"
  "strconv"
)

func characters(){
  fmt.Println(string( "Hello World"[0])) 
  var s string ="Hello"

  s = strings.Replace(s,"H", "X", 1)
  fmt.Println(s) 
  fmt.Println(strings.Contains(s, "lo"))

  fmt.Println(`""`) 
  fmt.Println("\"")

}

func bools(){
  t, f := true, false 
  fmt.Printf("%T %v %t\n", t, 1, t)
  fmt.Printf("%T %v %t\n", f, 0, f)
}

func type_convert(){ 
  var x int = 1
  xx := float64(x)

  fmt.Printf("%T %v %f\n", xx, xx, xx)

  var y float64 = 1.2
  yy := int(y)
  fmt.Printf("%T %v %d\n", yy, yy, yy)

  var s string = "14"
  i, _ := strconv.Atoi(s)
  fmt.Printf("%T %v\n", i, i)

  h := "Hello World"
  fmt.Println(string(h[0]))
}

func main(){
  var a [2]int
  a[0] = 1
  a[1] = 2
  fmt.Println(a)

  var b []int = []int{1, 2}
  b = append(b, 3) 
  fmt.Println(b)
}
