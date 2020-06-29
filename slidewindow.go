package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	st := os.Args[1:]
	if len(st) != 1 {
		fmt.Println("Wrong incoming. need only 1 varaiable")
	}
	fmt.Println(longestSubstring(st[0]))
}
func longestSubstring(st string) int {
	s := ""
	maxlen := 0
	curlen := 0
	index := 0
	for i := index; i < len(st); i++ {
		fmt.Println(s, string(st[i]))
		if strings.ContainsRune(s, rune(st[i])) {
			curlen = len(s)

			if curlen > maxlen {
				maxlen = curlen
			}
			index++
			i = index
			s = string(st[i])
		} else {
			s += string(st[i])
		}

	}
	if len(s) > maxlen {
		maxlen = len(s)
	}
	return maxlen
}
