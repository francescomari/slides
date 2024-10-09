package iterator4

import (
	"fmt"
	"iter"
)

// BEGIN BACKWARD OMIT
func backward(s []int) iter.Seq2[int, int] { // HL
	return func(yield func(int, int) bool) {
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
	for i, x := range backward(s) { // HL
		if x == sentinel {
			break // HL
		}
		fmt.Println(i, x)
	}
}

// END SOLUTION OMIT
