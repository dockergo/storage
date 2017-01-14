package qiniu

import (
	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/storage/adapter"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	domain    = "xxxx.com2.z0.glb.qiniucdn.com" // 指定私有空间的域名
	targeturl = "xxxx"                          // 指定需要抓取的文件的url
	movekey   = "movekey"
	copykey   = "yourcopykey"
)

type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type Qiniu struct {
	adapter.StorageAdapter
	config *config.Config
	client *kodo.Client
}

func New(config *config.Config) *Qiniu {
	c := new(Qiniu)
	c.config = config
	c.Name = "qiniu"
	conf.ACCESS_KEY = config.Storage.Qiniu.AccessKey
	conf.SECRET_KEY = config.Storage.Qiniu.SecretKey
	c.client = kodo.New(0, nil)
	return c
}
