package bins

import "time"

type Bin struct {
	Id       string
	Private  bool
	CreateAt time.Time
	Name     string
}

type BinList struct {
	Bins []Bin
}

func NewBin(id string, private bool, name string) *Bin {
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
