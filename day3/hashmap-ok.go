
package main

import "fmt"

func main() {
	myDict := map[string]int{
		"Google": 1990,
		"Amazon": 2000,
	}

	year, ok := myDict["Akhil"]

	if ok {
		fmt.Println("Exists")
		fmt.Println("Founded:", year)
	} else {
		fmt.Println("User Not found")
	}
}
