package main

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		f := x
		x, y = y, x+y
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
