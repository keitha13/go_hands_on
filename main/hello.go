package main

import (
	"fmt"
)

type Mydata struct {
	Name string
	Data []int
}

func (md Mydata) PrintData() {
	fmt.Println("*** Mydata ***")
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
	fmt.Println("*** end ***")
}

func main() {
	taro := Mydata{
		"hanako", []int{98, 76, 54, 32, 10},
	}
	taro.PrintData()
}
