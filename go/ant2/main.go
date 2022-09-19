package main

import "fmt"

func main() {
	for i := range ant(1_000_000) {
		fmt.Print(i)
	}
}
func ant(n int) <-chan int {
	c := gen(1)
	for i := 0; i < n; i++ {
		c = next(c)
	}
	return c
}

func gen(n int) <-chan int {
	c := make(chan int)
	go func() {
		c <- n
		close(c)
	}()
	return c
}

func next(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		prev, count := <-in, 1
		for value := range in {
			if value == prev {
				count++
			} else {
				out <- prev
				out <- count
				prev, count = value, 1
			}
		}
		out <- prev
		out <- count
		close(out)
	}()
	return out
}
