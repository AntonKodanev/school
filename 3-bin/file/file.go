package file

import (
	"encoding/json"
	"os"
)

type Files struct {
	Name string
	Path string
}

func NewFiles(name string, path string) *Files {
	return &Files{
		Name: name,
		Path: path,
	}
}

func (f *Files) ReadFile() ([]byte, error) {
	data, err := os.ReadFile(f.Path + "/" + f.Name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *Files) IsJsonFile() bool {
	data, _ := os.ReadFile(f.Path + "/" + f.Name)
	if json.Valid(data) {
		return true
	}
	return false
}
