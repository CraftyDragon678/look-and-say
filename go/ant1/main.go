package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "1"
	var max int
	fmt.Scanf("%d", &max)

	for i := 0; i < max; i++ {
		ch := s[0]
		count := 1
		next := ""
		for j := 1; j < len(s); j++ {
			if ch == s[j] {
				count += 1
			} else {
				next += string(ch) + strconv.Itoa(count)
				ch = s[j]
				count = 1
			}
		}
		next += string(ch) + strconv.Itoa(count)
		s = next
	}
	fmt.Println(s)
}
