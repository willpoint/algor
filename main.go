package main

import "fmt"

func solution(s string) int {
	max := 0
	seen := make(map[int]int)
	for i := range s {
		if num, balanced := isBalanced(s[i:]); balanced {
			seen[i] = num
		}
	}
	for i := range seen {
		if seen[i] > max {
			max = seen[i]
		}
	}
	return max
}

func isBalanced(s string) (int, bool) {
	i := make(map[string]int)
	f := string(s[0])
	n := ""
	for _, j := range s {
		l := string(j)
		if l != n {
			n = l
		}
		i[l]++
	}
	if len(i) == 2 {
		if (i[f]+i[n])%2 == 1 {
			return i[f] + i[n] - 1, true
		}
		return i[f] + i[n], true
	}
	return 0, false
}

func main() {
	j := "cabbacc"
	a := solution(j)
	fmt.Println(a)
}
