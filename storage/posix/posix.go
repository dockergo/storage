package posix

import (
	"os"
	"path/filepath"

	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/storage/adapter"
	"github.com/flyaways/storage/util/log"
)

type Posix struct {
	adapter.StorageAdapter
	Config *config.Config
}

func New(config *config.Config) *Posix {
	posix := new(Posix)
	posix.Config = config
	posix.Name = "posix"
	return posix
}

func (posix *Posix) IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func (posix *Posix) DirChecker(dir string) *errors.Error {
	if !posix.IsExist(dir) {
		log.Error("[%s:NoSuchBucket]", dir)
		return errors.NoSuchBucket
	}
	return nil
}

func (posix *Posix) FileChecker(filepath string) *errors.Error {
	if !posix.IsExist(filepath) {
		log.Error("[%s:NoSuchKey]", filepath)
		return errors.NoSuchKey
	}
	return nil
}

func (posix *Posix) getBucketPath(bucket string) string {
	return filepath.Join(posix.Config.Storage.Posix.Addr, bucket)
}

func walkDir(dirPth string) (files []os.FileInfo, filenames []string, err error) {
	files = make([]os.FileInfo, 0, 30)
	filenames = make([]string, 0, 30)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		fi.Name()
		files = append(files, fi)
		filenames = append(filenames, filename)
		return nil
	})

	return files, filenames, err
}
