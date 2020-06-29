package main

import "fmt"

func main() {
	n := 3999
	fmt.Println(intToRoman(n))
}
func intToRoman(n int) string {

	s := ""
	p := 1
	for n > 0 {
		s = convert(n%10, p) + s
		p++
		n /= 10
	}
	return s
}
func convert(a, b int) string {

	s := ""
	p := 1
	t := 0
	// rom := map[int]string{1: "I", 5: "V", 10: "X", 4: "IV", 9: "IX", 50: "L", 100: "C", 40: "XL", 90: "XC", 500: "D", 1000: "M", 400: "CD", 900: "CM"}
	for i := 1; i < b; i++ {
		p = p * 10
	}
	for a > 0 {
		// fmt.Println(a, p)
		if a == 4 || a == 5 || a == 9 {
			t = a
		} else {
			t = 1
		}
		switch t * p {
		case 1:
			s = "I" + s
		case 4:
			s = "IV" + s
		case 5:
			s = "V" + s
		case 9:
			s = "IX" + s
		case 10:
			s = "X" + s
		case 40:
			s = "XL" + s
		case 50:
			s = "L" + s
		case 90:
			s = "XC" + s
		case 100:
			s = "C" + s
		case 400:
			s = "CD" + s
		case 500:
			s = "D" + s
		case 900:
			s = "CM" + s
		case 1000:
			s = "M" + s
		}
		a = a - t
	}
	return s
}
