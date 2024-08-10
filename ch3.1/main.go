package main

import (
	"fmt"
	"math"
)

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

	var n int8 = 127

	fmt.Printf("%08b %08b\n ", n, n<<1) //01111111 -0000010

	// medals := []string{"gold", "silver", "bronze"}

	var i uint = 0
	fmt.Println(i - 1) //18446744073709551615

	/**
	panic: runtime error: index out of range [18446744073709551615] with length 3

	goroutine 1 [running]:
	main.main()
	*/
	// for i := uint(len(medals) - 1); i >= 0; i-- {
	// 	fmt.Println(medals[i], i) // "bronze", "silver", "gold"
	// }

	ascii := 'a'
	unicode := 's'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 Image 'Image'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
	var Avogadro float64 = 6.02214129e23
	fmt.Printf("%f\n", Avogadro)
	var v float64
	fmt.Println(math.IsNaN(v / v))

}
