package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//1.2 CLI Args
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	for i, arg := range os.Args[0:] {
		fmt.Println(i, arg)
	}

	t1 := time.Now()

	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(time.Since(t1))

	t2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(t2))

}
