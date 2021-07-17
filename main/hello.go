package main

import (
	"fmt"
)

type General interface{}

type GData interface {
	Set(nm string, g General) GData
	Print()
}

type NData struct {
	Name string
	Data int
}

func (nd *NData) Set(nm string, g General) GData {
	nd.Name = nm
	nd.Data = g.(int)
	return nd
}

func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

type SData struct {
	Name string
	Data string
}

func (sd *SData) Set(nm string, g General) GData {
	sd.Name = nm
	sd.Data = g.(string)
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s [%s] *\n", sd.Name, sd.Data)
	fmt.Println()
}

func main() {
	var data = []GData{}
	data = append(data, new(NData).Set("taro", 123))
	data = append(data, new(SData).Set("jiro", "hello"))
	data = append(data, new(NData).Set("saburo", 790))
	data = append(data, new(SData).Set("shiro", "great?"))
	for _, ob := range data {
		ob.Print()
	}
}
