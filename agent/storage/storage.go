package storage

import (
	"fmt"
	"strings"

	"github.com/flyaways/tracker"

	"github.com/flyaways/storage/agent/config"
	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/storage/kdfs"
	"github.com/flyaways/storage/agent/storage/nfs"
	"github.com/flyaways/storage/agent/storage/posix"
	"github.com/flyaways/storage/agent/storage/s3"
	"github.com/flyaways/storage/agent/storage/swift"
	"github.com/flyaways/storage/agent/util/log"
)

const (
	S3    = "s3"
	SWIFT = "swift"
	KDFS  = "kdfs"
	POSIX = "posix"
	NFS   = "nfs"
)

func NewStorage(config *config.Config) (adapter.Storager, error) {
	log.Info("[storage-type:%s]", tracker.Red(config.Storage.Type))
	switch strings.ToLower(config.Storage.Type) {
	case SWIFT:
		return swift.New(config), nil
	case S3:
		return s3.New(config), nil
	case POSIX:
		return posix.New(config), nil
	case KDFS:
		return kdfs.New(config), nil
	case NFS:
		return nfs.New(config), nil
	default:
		log.Error(tracker.Red("[unsupport storage type]"))
		return nil, fmt.Errorf(tracker.Red("[storage type error]"))
	}
}
