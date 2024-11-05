package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 50
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))
}
