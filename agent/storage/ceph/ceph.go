package ceph

import (
	"time"

	"wpsep.net/storage-agent/agent/config"

	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

type Ceph struct {
	adapter.StorageAdapter
	config *config.Config
	client *s3.S3
}

func New(config *config.Config) *Ceph {
	c := new(Ceph)
	c.config = config
	c.Name = "ceph"
	auth, err := aws.GetAuth(config.Storage.Ceph.AccessKey, config.Storage.Ceph.SecretKey, "", time.Time{})
	if err != nil {
		panic(err)
	}
	c.client = s3.New(auth, aws.Region{
		Name:              "cn-north-1",
		S3Endpoint:        config.Storage.Ceph.Addr,
		S3LowercaseBucket: true,
	})

	return c
}
