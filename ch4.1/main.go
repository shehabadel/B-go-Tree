package main

import (
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

	for i, v := range symbol {
		log.Default().Println(i, v)
	}

	r := [...]int{99: -1}
	for _, v := range r {
		fmt.Printf("%d \t", v)
	}
}
