package main

import "fmt"

func main() {
	st := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(st))
}
func longestCommonPrefix(st []string) string {
	s := ""
	in := 0
	if len(st) == 1 {
		return st[0]
	}
	if len(st) == 2 {
		if len(st[0]) > len(st[1]) {
			st[0], st[1] = st[1], st[0]
		}
		for i := 0; i < len(st[0]); i++ {
			if st[0][i] != st[1][i] {
				return s
			}
			s += string(st[0][i])
		}
	} else {
		for i := 0; i < len(st)-1; i++ {

			fmt.Println(i, in, s)
			if !valid(in, st) {
				// fmt.Println("tut")
				return s
			}
			if st[i][in] != st[i+1][in] {
				// fmt.Println("tut 1", s)
				return s
			}
			if i == len(st)-2 {
				s += string(st[i][in])
				// fmt.Println(s)
				i = -1
				in = in + 1
			}

		}
	}

	return s
}
func valid(a int, s []string) bool {
	for _, k := range s {
		if a >= len(k) {
			return false
		}
	}
	return true
}
