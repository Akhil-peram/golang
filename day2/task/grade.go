package main

import "fmt"

func main(){
	grade:=92
	switch {
		case grade>=90:
		fmt.Println("A")
		case grade>=80:
		fmt.Println("B")
		case grade>=70:
		fmt.Println("C")
		case grade>=60:
		fmt.Println("D")
		default:
		fmt.Println("Failed")
	}
}
