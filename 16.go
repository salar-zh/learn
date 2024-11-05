package main

import (
	"fmt"
	"reflect"
)

func main() {
	MyFloatNumber := 6.9
	fmt.Println("MyFloatNumber:", MyFloatNumber)
	MyIntNumber := int(MyFloatNumber)
	fmt.Println("MyIntNumber:", MyIntNumber)
	fmt.Println(reflect.TypeOf(MyIntNumber))
}
