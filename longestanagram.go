package main

import (
	"fmt"
	"os"
)

func main() {
	st := os.Args[1:]
	if len(st) != 1 {
		fmt.Println("                 IDI NA X**")
		return
	}
	fmt.Println(longestAnagram(st[0]))
}
func longestAnagram(a string) string {
	curstr := ""
	polstr := string(a[0])
	in := 0
	for i := len(a) - 1; i > in; i-- {
		fmt.Println(i, in)
		if a[in] == a[i] {
			curstr = a[in : i+1]
			if isitPolindrome(curstr) && len(curstr) > len(polstr) {
				polstr = curstr
			}
		}
		if i == in+1 {
			in++
			i = len(a)
		}
	}
	return polstr
}
func isitPolindrome(s string) bool {
	if len(s)%2 != 0 {
		for i := 0; i < len(s)/2-1; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
