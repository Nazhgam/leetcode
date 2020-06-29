package main

import "fmt"

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
func romanToInt(s string) int {
	n := 0
	i := 0
	for i < len(s) {

		if s[i] == 'I' {
			if i < len(s)-1 && s[i+1] == 'V' {
				n = n + 4
				i = i + 2
			} else if i < len(s)-1 && s[i+1] == 'X' {
				n = n + 9
				i = i + 2
			} else {
				n = n + 1
				i = i + 1
			}
		} else if s[i] == 'X' {
			if i < len(s)-1 && s[i+1] == 'L' {
				n = n + 40
				i = i + 2
			} else if i < len(s)-1 && s[i+1] == 'C' {
				n = n + 90
				i = i + 2
			} else {
				n = n + 10
				i = i + 1
			}
		} else if s[i] == 'C' {
			if i < len(s)-1 && s[i+1] == 'D' {
				n = n + 400
				i = i + 2
			} else if i < len(s)-1 && s[i+1] == 'M' {
				n = n + 900
				i = i + 2
			} else {
				n = n + 100
				i = i + 1
			}
		} else if s[i] == 'V' {
			n = n + 5
			i++
		} else if s[i] == 'L' {
			n = n + 50
			i++
		} else if s[i] == 'D' {
			n = n + 500
			i++
		} else if s[i] == 'M' {
			n = n + 1000
			i++
		}
	}
	return n
}
