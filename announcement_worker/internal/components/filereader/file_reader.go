package filereader

import "os"

type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

type filereader struct {
}

func NewFileReader() FileReader {
	return &filereader{}
}

func (f *filereader) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
