package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world👉"

	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len(s))

	fmt.Println(s[0:len(s)])
	fmt.Println(s[0:utf8.RuneCountInString(s)]) // mesh httl3 el emoji mazbot 3ashan its not an ascii code point, a substring requires the bytes of original string, not the count of characters

	fmt.Printf("%08b\n", []byte("a")) // A high-order 0 indicates 7-bit ASCII
	fmt.Printf("%08b\n", []byte("👉")) // higher-order 1111 indicates 4 bytes sequence
	/**
	  1st byte (11110000):
	  Starts with 11110, indicating a 4-byte sequence
	  The remaining 000 are part of the code point

	  2nd byte (10011111):
	  Starts with 10, as required for continuation bytes
	  011111 contributes to the code point

	  3rd byte (10010001):
	  Starts with 10
	  010001 contributes to the code point

	  4th byte (10001001):
	  Starts with 10
	  001001 contributes to the code point
	*/

}
