package posix

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util/log"
)

func (posix *Posix) InitBucket(initdata *adapter.InitData) error {
	err := os.MkdirAll(posix.getBucketPath(initdata.Bucket), os.ModePerm)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		return err
	}

	/***
		if err := posix.Checker.DirChecker(posix.getBucketPath(initdata.Bucket)); err != nil {
			log.Error("[%s:%s]", posix.Name, err.Error())
			return err
		}
	***/
	return nil
}

func (posix *Posix) InitObject(initdata *adapter.InitData) error {

	if err := os.MkdirAll(posix.getBucketPath(initdata.Bucket), os.ModePerm); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		return err
	}

	/***
		if err := posix.Checker.DirChecker(posix.getBucketPath(initdata.Bucket)); err != nil {
			log.Error("[%s:%s]", posix.Name, err.Error())
			return err
		}
	***/
	filename := filepath.Join(posix.getBucketPath(initdata.Bucket), initdata.Key)
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		return err
	}

	if err := ioutil.WriteFile(filename, initdata.RawData, 0666); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		return err
	}

	return nil
}
