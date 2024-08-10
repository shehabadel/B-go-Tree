package main

import "fmt"

var i int = 1

var j int32 = 2

func main() {
	// j=i; //compiler error

	//requires conversion
	j = int32(i)

	var a, b int = 1, 3
	//  01   10
	// 	01	 01

	// 10	01
	// 10	10

	// 01 	11
	// 01	00
	c := a &^ b

	fmt.Println(c)

	var z uint8 = 1 << 1
	var x uint8 = z | 1<<5
	var y uint8 = z | 1<<2

	fmt.Printf("%08b \n", z)
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
}
