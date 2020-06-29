package main

import "fmt"

func main() {
	fmt.Println(pow(2, 72))
}
func pow(a, b int64) int64 {
	if b == 0 {
		return 1
	}

	if b == 1 {
		return a
	}
	return a * pow(a, b-1)
}
