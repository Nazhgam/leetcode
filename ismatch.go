package 

import "fmt"

func main() {
	var s, p string
	s = "a"
	p = "a**"
	fmt.Println(isMatch(s, p))
}

func isMatch(s, p string) bool {
	
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 || p[1] != '*' {
		if len(s) < 1 || (p[0] != '.' && s[0] != p[0]) {
			return false
		}
		return isMatch(s[1:], p[1:])
	} else {
		l := len(s)
		i := -1
		for i < l && (i < 0 || p[0] == '.' || p[0] == s[i]) {
			fmt.Println(s, p,i)
			if isMatch(s[i+1:], p[2:]) {
				return true

			}
			i++
		}
	}

	return false
}
