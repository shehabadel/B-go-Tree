package main

import (
	"b-go-tree/ch2.3/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC) // "100" °C
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.BoilingC)) // "180" °F

	var c tempconv.Celsius
	var f tempconv.Fahrenheit

	// fmt.Println(c == f) //compiler error type mismatch
	fmt.Println(c == tempconv.Celsius(f))
}
