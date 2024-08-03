package main

import (
	"b-go-tree/ch2.3/tempconv"
	"fmt"
	"os"
	"strconv"
)

var a = b + c // a initialized third, to 3
var b = f()   // b initialized second, to 2, by calling f
var c = 1     // c initialized first, to 1

func f() int { return c + 1 }
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Println(f, c, tempconv.FToC(f), tempconv.CToF(c))
	}
}
