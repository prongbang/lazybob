package lazybob

import (
	"io/ioutil"
)

type Loader interface {
	Load(filename string) ([]byte, error)
}

type fileLoader struct {
}

func (f *fileLoader) Load(filename string) ([]byte, error) {
	pemData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return pemData, err
}

func NewFileLoader() Loader {
	return &fileLoader{}
}
