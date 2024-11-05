package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 63.4
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))
}
