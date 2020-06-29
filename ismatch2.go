package main

import "fmt"

func main() {
	var s, p string
	s = "magga"
	p = ".*ga"
	fmt.Println(isMatch2(s, p))
}
func isMatch2(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 {
		if len(s) < 1 {
			return false
		}
		if s[0] != p[0] && p[0] != '.' {
			return false
		}

		return isMatch2(s[1:], p[1:])

	}
	if p[1] != '*' {
		if len(s) < 1 {
			return false
		}
		if p[0] != s[0] && p[0] != '.' {
			return false
		}
		return isMatch2(s[1:], p[1:])
	} else {
		if isMatch2(s, p[2:]) {
			return true
		}
		i := 0
		for i < len(s) && (s[i] == p[0] || p[0] == '.') {
			if isMatch2(s[i+1:], p[2:]) {
				return true
			}
			i++
		}
	}
	return false
}
