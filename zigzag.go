package main

import (
	"fmt"
	"os"
)

func main() {
	st := os.Args[1:]
	if len(st) != 1 {
		fmt.Println("need only one word\nyou stupid donkey")
		return
	}
	fmt.Println(zigzag(st[0], 4))
	return
}
func zigzag(a string, num int) string {
	s := ""
	curin := 0
	predin := 0
	arr := make(map[int]string)
	for i := 0; i < len(a); i++ {
		arr[curin] += string(a[i])
		if curin == num-1 {
			predin = curin
		} else if curin == 0 {
			predin = curin
		}
		if predin == 0 {
			curin++
		}
		if predin == num-1 {
			curin--
		}

	}
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}
