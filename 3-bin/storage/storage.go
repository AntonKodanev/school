package storage

import "os"

type StorageDB struct {
	filename string
}

func NewStorage() *StorageDB {
	return &StorageDB{
		filename: "storage.json",
	}
}

func (s *StorageDB) Read() ([]byte, error) {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *StorageDB) Write(data []byte) error {
	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
