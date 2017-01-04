package nfs

import (
	"fmt"
	"os"

	"github.com/flyaways/storage/agent/config"
	"github.com/flyaways/storage/agent/errors"
	"github.com/flyaways/storage/agent/storage/posix"
	execcmd "github.com/flyaways/storage/agent/storage/posix/exec"
	"github.com/flyaways/storage/agent/util/log"
)

type NFS struct {
	posix.Posix
}

func New(config *config.Config) *NFS {
	nfs := new(NFS)
	nfs.Config = config
	nfs.Name = "nfs"
	nfs.Checker = nfs
	return nfs
}

func (nfs *NFS) DirChecker(dir string) *errors.Error {
	if !nfs.IsExist(dir) {
		errCure := nfs.cure()
		if errCure != nil {
			return errors.NoSuchBucket
		}
	}
	return nil
}

func (nfs *NFS) FileChecker(filename string) *errors.Error {
	if !nfs.IsExist(filename) {
		return errors.NoSuchKey
	}
	return nil
}

func (nfs *NFS) cure() error {
	defer func() {
		if r := recover(); r != nil {
			log.Error("[E]", r)
		}
	}()

	localPath := nfs.Config.Storage.Nfs.Addr
	remotePath := nfs.Config.Storage.Nfs.RemoteAddr
	cmd := fmt.Sprintf(" mount -t nfs -o rw %s %s", remotePath, localPath)
	log.Info("mount cmd: %s", cmd)
	execcmd.Exec("/bin/sh", "-c", cmd)

	if stat := nfs.IsExist(localPath); !stat {
		return errors.NoSuchBucket
	} else {
		if denied := permission(localPath); denied != nil {
			return errors.PermissionDenied
		}
	}
	return nil
}

func permission(path string) error {
	defer func() {
		if r := recover(); r != nil {
			log.Error("[E]", r)
		}
	}()

	filename := fmt.Sprintf("%s/dhv5h5jdh7fjyeh97ydvg25evb22v342djjd", path)

	if mkdir := os.MkdirAll(filename, os.ModePerm); mkdir != nil {
		return mkdir
	}

	if _, createrr := os.Create(filename); createrr != nil {
		return createrr
	}

	if removerr := os.Remove(filename); removerr != nil {
		return removerr
	}

	return nil
}
