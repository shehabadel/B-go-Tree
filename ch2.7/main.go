package main

import "fmt"

var i int = 3

func f() {
	fmt.Println(i)
}
func main() {
	var i float64 = 1.0
	f()
	fmt.Println(i)
}
