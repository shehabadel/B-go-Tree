package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	s := make([]int, len(a))
	copy(s, a)
	reverse(s)
	fmt.Println(s)

	fmt.Println(a)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
