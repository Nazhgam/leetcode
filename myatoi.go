package main

import (
	"fmt"
	"os"
)

func main() {
	st := os.Args[1:]
	if len(st) != 1 {
		return
	}
	fmt.Println(myAtoi(st[0]))
}
func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	t := 0
	str = cut(str)
	str = checkzero(str)
	fmt.Println(str)
	if str == "" {
		return 0
	}

	if str[0] == '-' {
		if len(str) > 11 {
			return -2147483648
		}
		if len(str) == 11 && str > "-2147483648" {
			return -2147483648
		}
	}
	if str[0] == '+' {
		if len(str) > 11 {
			return 2147483647
		}
		if len(str) == 11 && str > "+2147483647" {
			return 2147483647
		}
	}
	if len(str) > 10 {
		return 2147483647
	}
	if len(str) == 10 && str > "2147483647" {
		return 2147483647
	}

	if str[0] == '-' {
		for _, k := range str[1:] {
			t = t*10 + int(k-48)
		}
		t *= -1
	} else {
		if str[0] == '+' {
			str = str[1:]
		}
		for _, k := range str {
			t = t*10 + int(k-48)
		}
	}

	return t
}
func checkzero(str string) string {
	if str == "" {
		return ""
	}
	if str[0] == '0' {
		return checkzero(str[1:])
	}
	if str[0] == '-' {
		return "-" + checkzero(str[1:])
	}
	if str[0] == '+' {
		return "+" + checkzero(str[1:])
	}
	return str
}
func cut(str string) string {
	if str == "" {
		return ""
	}
	if str[0] == ' ' {
		return cut(str[1:])
	}
	if str[len(str)-1] < '0' || str[len(str)-1] > '9' {
		return cut(str[:len(str)-1])
	}
	s := ""
	for i, k := range str {
		if k >= '0' && k <= '9' {
			s += string(k)
		} else if i == 0 && (k == '-' || k == '+') {
			s += string(k)
		} else {
			break
		}
	}
	return s
}
