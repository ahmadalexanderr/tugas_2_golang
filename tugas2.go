package main

import "fmt"

func main() {
	x := 5
	if x%2 == 0 {
		fmt.Println(x, "adalah bilangan genap")
	} else {
		fmt.Println(x, "adalah bilangan ganjil")
	}
	return
}
