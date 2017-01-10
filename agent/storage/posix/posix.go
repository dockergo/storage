package posix

import (
	"os"
	"path/filepath"

	"github.com/flyaways/storage/agent/config"
	"github.com/flyaways/storage/agent/errors"
	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util/log"
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
