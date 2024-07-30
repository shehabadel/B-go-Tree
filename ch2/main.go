package main

import "fmt"

func main() {
	var s string
	fmt.Println(s) // Prints ""

	var sa []string
	fmt.Println(sa)      //prints []
	fmt.Println(len(sa)) //prints 0

	// var i,j,k int // int,int,int
	// var b,f,d = true, 2.3, "four" //bool, float64, string

	// var f,err = os.Open(name) -> some fn that returns two variables

	// Short-variable declaration

	/**
	Takes the current form

	name := expression
	type of `name` defined by type of the `expression`

	*/
	t := 2.3 //Float64

	/**
	short variable declaration tends to be used for local variables
	A var declaration tends to be reserved for local variables that need an explicit type
	that differs from that of the initializer expression

	*/

	i := 100 //int

	var boiling float64 = 100 //float64

	j, k := 0, 1
	fmt.Println(t, i, boiling, j, k)

	i, j = j, i //Swapping i and j
	fmt.Println(t, i, boiling, j, k)
}
