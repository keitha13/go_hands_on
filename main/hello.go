package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

type Mydata struct {
	Name string
	Data []int
}

func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

type Yourdata struct {
	Name string
	Mail string
	Age  int
}

func (md *Yourdata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	md.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	md.Age = n
}

func (md *Yourdata) PrintData() {
	fmt.Printf("Im %s. (%d).\n", md.Name, md.Age)
	fmt.Printf("mail: %s.\n", md.Mail)
}

func main() {
	ob := [2]Data{}
	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{
		"name": "hayashi",
		"data": "33, 44, 55",
	})
	ob[1] = new(Yourdata)
	ob[1].SetValue(map[string]string{
		"name": "keita",
		"mail": "keita@mail.com",
		"age":  "66",
	})
	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}
