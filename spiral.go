package main

import "fmt"

func main() {
	var n int = 100

	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	t := n * n
	massive := make([]int, t)
	for i := 1; i <= t; i++ {
		massive[i-1] = i
	}

	c := 0
	riht := false
	left := false
	up := false
	down := false
	maxI := n - 1
	maxJ := n - 1
	minI := 0
	minJ := 0
	i := 0
	j := 0
	for {
		if c == t-1 {
			arr[i][j] = massive[c]
			break
		}
		if i == minI && j != maxJ {
			// for j < maxJ {

			arr[i][j] = massive[c]
			c++
			j++
			// }

			riht = true
			if up == true {
				up = false
				minJ += 1
			}
		}
		if j == maxJ && i != maxI {
			// for i < maxI {
			arr[i][j] = massive[c]
			c++
			i++
			// }
			if riht == true {
				minI += 1
				riht = false
			}
			down = true
		}
		if i == maxI && j != minJ {
			// for j > minJ {
			arr[i][j] = massive[c]
			c++
			j--
			// }
			if down == true {
				down = false
				maxJ -= 1
			}
			left = true
		}
		if j == minJ && i != minI {
			// for i > minI {
			arr[i][j] = massive[c]
			c++
			i--
			// }
			if left == true {
				left = false
				maxI -= 1
			}
			up = true
		}
	}
	for _, k := range arr {
		fmt.Println(k)
	}
	// fmt.Println(arr, massive, i, j, c, maxI, maxJ, minI, minJ, riht, down, up, left)
}
