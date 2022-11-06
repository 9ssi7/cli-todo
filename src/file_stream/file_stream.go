package file_stream

import (
	"encoding/json"
	"errors"
	"os"
)

type FileStream struct {}

func New() *FileStream {
	return &FileStream{}
}

func (fs *FileStream) Read(fn string, e interface{}) error {
	file, err := os.ReadFile(fn)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if err = json.Unmarshal(file, e); err != nil {
		return err
	}
	return nil
}

func (fs *FileStream) Write(fn string, d interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	return os.WriteFile(fn, data, 0644)
}