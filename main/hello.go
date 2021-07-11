package main

import (
	"fmt"
)

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := new(Mydata)
	fmt.Println(taro)
	taro.Name = "taro"
	taro.Data = make([]int, 5)
	fmt.Println(taro)
}
