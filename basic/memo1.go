package main
import (
  "fmt"
  "os/user"
  "time"
)

var (
  i int = 1
  f64 float64 = 1.2
  s string = "test"
  // var t, f bool = true, false
  t, f bool = true, false                       // t and f output false

// init function
func init(){
    fmt.Println("Init")
}

func variable(){
  /* Same as belows
  var i int =1
  var f64 float64 = 1.2
  var s string = "test"

  var (
    ii int =1
    ff64 float64 = 1.2
    ss string = "test"
    // var tt, ff bool = true, false
    tt, ff bool                       // t and f output false
  )
  fmt.Println(ii,ff64, ss, tt, ff)
  */
}

// main function
func main(){
    variable()
    fmt.Println("Hello", "World")
    fmt.Println(time.Now())
    fmt.Println(user.Current())
}

// how to use godoc
// $ go doc command
