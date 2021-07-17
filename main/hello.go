package main

import (
	"fmt"
)

type General interface{}

type GData interface {
	Set(nm string, g General)
	Print()
}

type GDataImpl struct {
	Name string
	Data General
}

func (gd *GDataImpl) Set(nm string, g General) {
	gd.Name = nm
	gd.Data = g
}

func (gd *GDataImpl) Print() {
	fmt.Printf("<<%s>>", gd.Name)
	fmt.Println(gd.Data)
}

func main() {
	var data = []GDataImpl{}
	data = append(data, GDataImpl{"keita", 123})
	data = append(data, GDataImpl{"hayashi", "hello"})
	data = append(data, GDataImpl{"hayata", []int{123, 456, 789}})
	for _, ob := range data {
		ob.Print()
	}
}
