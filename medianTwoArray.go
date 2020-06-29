package main

import "fmt"

func main() {
	arr1 := []int{}
	arr2 := []int{3, 4}
	arr1 = sortarr(arr1)
	arr2 = sortarr(arr2)

	fmt.Println(mediantwoarr(arr1, arr2))
}
func sortarr(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
func mediantwoarr(a, b []int) float64 {
	var merge []int

	for {

		if len(a) == 0 && len(b) == 0 {
			if len(merge)%2 != 0 {
				return float64(merge[len(merge)/2])
			} else {
			
				return (float64(merge[len(merge)/2]+merge[len(merge)/2-1]) / 2)
			}
		}
		if len(a) == 0 {
			for _, k := range b {
				merge = append(merge, k)
			}
			b = a
		}
		if len(b) == 0 {
			for _, k := range a {
				merge = append(merge, k)
			}
			a = b
		}

		if len(a) != 0 && len(b) != 0 {
			if a[0] > b[0] {
				a, b = b, a
			}
			j := startandendposition(a, b)

			for i := 0; i <= j; i++ {
				merge = append(merge, a[i])
			}
			merge = append(merge, b[0])
			a = a[j+1:]
			b = b[1:]
		}
	}

	return 1.0
}
func startandendposition(a, b []int) int {
	var j int
	if a[0] > b[0] {
		a, b = b, a
	}
	for index, k := range a {
		if k > b[0] {
			j = index - 1
			return j
		}
	}
	j = len(a) - 1
	return j
}
