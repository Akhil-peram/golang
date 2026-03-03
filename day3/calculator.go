package main

import "fmt"

func main(){
  var value bool = true
  for{
    fmt.Println("Enter num1 and num2 : ")
    var x,y int = 0,0
    fmt.Scan(&x,&y)
    fmt.Println(x+y)
}
