package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x *int
	var y int = 10
	x = &y
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(&y))
	fmt.Println(reflect.TypeOf(x) == reflect.TypeOf(&y))
	fmt.Println(&y)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*x)

	fmt.Println(f()) // "false"
	fmt.Println(f()) // "false"
	fmt.Println(f()) // "false"
	fmt.Println("\n\n\n\n\n")

	p := new(int)
	fmt.Println(p)
	fmt.Println(*p)
}

func f() *int {
	v := 1
	return &v
}
