package main

import (
	"fmt"
	"time"
)

type Bin struct {
	Id       string
	Private  bool
	CreateAt time.Time
	Name     string
}

type BinList struct {
	Bins []Bin
}

func newBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:       id,
		Private:  private,
		CreateAt: time.Now(),
		Name:     name,
	}

}

func NewBinList() *BinList {
	return &BinList{
		Bins: []Bin{},
	}
}

func main() {
	bin := newBin("1", true, "test")
	fmt.Println(bin)
	binList := NewBinList()
	fmt.Println(binList)
}
