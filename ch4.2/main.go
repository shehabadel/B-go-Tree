package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	s := make([]int, len(a))
	copy(s, a)
	reverse(s)
	fmt.Println(s)

	fmt.Println(a)

	var q []int
	g := []int{1, 2, 3}
	g = append(g, 4, 5, 6, 7)
	fmt.Println(q == nil, g == nil, cap(g), g)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		fmt.Println("before: ", x, cap(x), len(x))
		x = y
		fmt.Println("after: ", x, cap(x), len(x))

	}

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
