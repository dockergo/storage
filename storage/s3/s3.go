package s3

import (
	"github.com/flyaways/storage/config"

	"github.com/flyaways/storage/storage/adapter"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

type S3c struct {
	adapter.StorageAdapter
	config *config.Config
	client *s3.S3
}

func New(config *config.Config) *S3c {
	c := new(S3c)
	c.config = config
	c.Name = "s3"

	auth, err := aws.GetAuth(config.Storage.S3.AccessKey, config.Storage.S3.SecretKey)
	if err != nil {
		panic(err)
	}

	var cnc = aws.Region{
		Name:                 "cn-north-1",
		S3Endpoint:           c.config.Storage.S3.Addr,
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
	}

	c.client = s3.New(auth, cnc)

	return c
}
