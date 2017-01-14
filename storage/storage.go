package storage

import (
	"fmt"
	"strings"

	"github.com/flyaways/tracker"

	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/storage/adapter"
	"github.com/flyaways/storage/storage/oss"
	"github.com/flyaways/storage/storage/posix"
	"github.com/flyaways/storage/storage/qiniu"
	"github.com/flyaways/storage/storage/s3"
	"github.com/flyaways/storage/storage/swift"
	"github.com/flyaways/storage/util/log"
)

const (
	S3    = "s3"
	SWIFT = "swift"
	POSIX = "posix"
	OSS   = "oss"
	QINIU = "qiniu"
)

func NewStorage(config *config.Config) (adapter.Storager, error) {
	log.Info("[storage-type:%s]", tracker.Yellow(config.Storage.Type))
	switch strings.ToLower(config.Storage.Type) {
	case SWIFT:
		return swift.New(config), nil
	case S3:
		return s3.New(config), nil
	case POSIX:
		return posix.New(config), nil
	case OSS:
		return oss.New(config), nil
	case QINIU:
		return qiniu.New(config), nil
	default:
		log.Error(tracker.Red("[unsupport storage type]"))
		return nil, fmt.Errorf(tracker.Red("[storage type error]"))
	}
}
