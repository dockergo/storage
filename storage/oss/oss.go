package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/storage/adapter"
	"github.com/flyaways/storage/util/log"
)

type OSS struct {
	adapter.StorageAdapter
	config *config.Config
	client *oss.Client
}

func New(config *config.Config) *OSS {
	ossc := new(OSS)
	ossc.config = config
	ossc.Name = "oss"
	client, err := oss.New(config.Storage.OSS.Addr, config.Storage.OSS.AccessKey, config.Storage.OSS.SecretKey)
	ossc.client = client
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		panic("oss auth failed")
	}
	return ossc
}
