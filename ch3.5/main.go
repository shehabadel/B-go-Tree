package main

import (
	"bytes"
	"fmt"
	"strings"
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

	a := basename("/a/v/b/c/d/e.d.go")
	fmt.Println(a)
	v := comma("1923234223")
	fmt.Println(v)
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
	fmt.Println(commaIteration("1234"))

}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := utf8.RuneCountInString(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func commaIteration(s string) string {
	var buf bytes.Buffer

	var pre int
	pre = len(s) % 3
	if pre == 0 {
		pre = 3
	}

	buf.WriteString(s[:pre])

	for i := pre; i < len(s); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
