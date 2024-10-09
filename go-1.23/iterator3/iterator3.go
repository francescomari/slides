package iterator3

import "fmt"

// BEGIN BACKWARD OMIT
func backward(s []int) func(func(int, int) bool) { // HL
	return func(yield func(int, int) bool) { // HL
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

// END BACKWARD OMIT

// BEGIN SOLUTION OMIT
func solution(s []int, sentinel int) {
	forEachElement := backward(s) // HL

	forEachElement(func(i, x int) bool { // HL
		if x == sentinel {
			return false
		}
		fmt.Println(i, x)
		return true
	})
}

// END SOLUTION OMIT
