package main

import "fmt"

const Pi = 3.14

const (
  Username = "name"
  Password = "pass"
)

const big = 123456789

func main(){
  fmt.Println(Pi, Username, Password)
  fmt.Println(big) 

  var (
    u8 uint8 = 255
    i8 int8 = 127
    f32 float32 = 0.2
    c64 complex64 = -5 + 12i
  )
  fmt.Println(u8, i8, f32, c64)

  fmt.Printf("type=%T value=%v",u8, u8)

  fmt.Println()
}
