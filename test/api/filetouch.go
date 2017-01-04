package api_test

import (
	"io/ioutil"
	"os"
)

func Exist(filename string) bool {
	_, err := os.Stat(*curfile)
	return err == nil || os.IsExist(err)
}

func TouchFile() error {
	if !Exist(*curfile) {
		return ioutil.WriteFile(*curfile, []byte(*content), 0666)
	}
	return nil
}
