package main

import "fmt"

func main(){
  nums := []int{1,2,4,5,6,,8,12,16}

  for i ,v:= range nums{
    fmt.Println(i,v)
}
