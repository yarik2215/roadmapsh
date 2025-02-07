package tracker

import (
	"encoding/json"
	"io"
	"os"
)

type FileStorage struct {
	fileName string
}

func asserError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewFileStorage(fileName string) FileStorage {
	return FileStorage{fileName: fileName}
}

func (s FileStorage) Load(dst *[]Task) {
	f, err := os.Open(s.fileName)
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(s.fileName)
	}
	asserError(err)
	defer f.Close()

	data, err := io.ReadAll(f)
	asserError(err)
	json.Unmarshal(data, dst)
}

func (s FileStorage) Save(src []Task) {
	f, err := os.Create(s.fileName)
	asserError(err)
	defer f.Close()

	err = json.NewEncoder(io.NewOffsetWriter(f, 0)).Encode(src)
	asserError(err)
}
