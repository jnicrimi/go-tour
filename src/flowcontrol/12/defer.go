package main

import "fmt"

func main() {
	defer println("world")
	fmt.Println("hello")
}
