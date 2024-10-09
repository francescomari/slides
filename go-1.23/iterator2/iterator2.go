package iterator2

import "fmt"

// BEGIN SOLUTION OMIT
func backward(s []int, yield func(int, int) bool) { // HL
	for i := len(s) - 1; i >= 0; i-- {
		if !yield(i, s[i]) { // HL
			return
		}
	}
}

func solution(s []int, sentinel int) {
	backward(s, func(i, x int) bool { // HL
		if x == sentinel {
			return false
		}
		fmt.Println(i, x)
		return true
	})
}

// END SOLUTION OMIT
