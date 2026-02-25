package main

import (
	"demo/bin/bins"
	"fmt"
)

func main() {
	bin := bins.NewBin("1", true, "test")
	fmt.Println(bin)
	binList := bins.NewBinList()
	fmt.Println(binList)
}
