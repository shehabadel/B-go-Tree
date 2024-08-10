package main

import "fmt"

var i int = 3

func f() {
	fmt.Println(i)
}

func k() int {
	return 1
}

func g(x int) int {
	return 1
}
func main() {
	var i float64 = 1.0
	f()
	fmt.Println(i)

	if x := k(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	// fmt.Println(x, y) //compile error, no access to x and y

	// x and y scopes are limited to their if statements only
}
