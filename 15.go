package main

import "fmt"

func main() {
	a := 2
	b := 4
	fmt.Println((a = b) || (b == a))
}
