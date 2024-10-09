package iterator1

import "fmt"

// BEGIN SOLUTION OMIT
func solution(s []int, sentinel int) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == sentinel {
			break
		}
		fmt.Println(i, s[i])
	}
}

// END SOLUTION OMIT
