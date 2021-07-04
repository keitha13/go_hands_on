package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30, 40, 50, 60}
	b := a[1:3]
	fmt.Println(a)
	fmt.Println(b)
}
