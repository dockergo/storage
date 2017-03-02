package types

import (
	"fmt"
	"strings"

	"github.com/flyaways/tracker"

	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/types/adapter"
	"github.com/flyaways/storage/types/oss"
	"github.com/flyaways/storage/types/posix"
	"github.com/flyaways/storage/types/qiniu"
	"github.com/flyaways/storage/types/s3"
	"github.com/flyaways/storage/types/swift"
	"github.com/flyaways/storage/util/log"
)

const (
	S3    = "s3"
	SWIFT = "swift"
	POSIX = "posix"
	OSS   = "oss"
	QINIU = "qiniu"
)

func New(config *config.Config) (adapter.Storager, error) {
	log.Info(tracker.Blue("[storage-type:%s]", config.Storage.Type))
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
		log.Error(tracker.Red("[unsupport type]"))
		return nil, fmt.Errorf(tracker.Red("[type error]"))
	}
}
