package main

import "fmt"

func main() {
	ar := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(ar))
}
func maxArea(ar []int) int {
	s := 0
	e := len(ar) - 1
	a := 0
	for s <= e {
		if ar[s] > ar[e] {
			if ar[e]*(e-s) > a {
				a = ar[e] * (e - s)
			}
			e--
		}
		if ar[s] <= ar[e] {

			if ar[s]*(e-s) > a {
				a = ar[s] * (e - s)
			}
			s++
		}
	}

	return a
}
