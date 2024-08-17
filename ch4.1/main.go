package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
)

func main() {
	const (
		USD = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", "€", GBP: "£", RMB: "¥"}

	f := flag.String("size", "256", "size of bit words")

	flag.Parse()
	for i, v := range symbol {
		log.Default().Println(i, v)
	}

	fmt.Println([]uint8("x"), []byte("x"), []string{"x"})

	var c1 []byte

	fmt.Println(*f)
	switch *f {
	case "256":
		sum := sha256.Sum256([]byte("x"))
		c1 = sum[:]
	case "512":
		sum := sha512.Sum512([]byte("x"))
		c1 = sum[:]
	case "384":
		sum := sha512.Sum384([]byte("x"))
		c1 = sum[:]
	}
	log.Default().Printf("%[1]x \t %[1]s", c1)

}

func setZero(ptr *[4]int) {
	for i := range ptr {
		ptr[i] = 0
	}
}
