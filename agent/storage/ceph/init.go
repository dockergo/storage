package ceph

import (
	"github.com/goamz/goamz/s3"
	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util/log"
)

func (c *Ceph) InitBucket(initdata *adapter.InitData) error {
	if err := c.client.Bucket(initdata.Bucket).PutBucket(s3.PublicRead); err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		return err
	}

	return nil
}

func (c *Ceph) InitObject(initdata *adapter.InitData) error {
	conType := "application/ocet-stream"
	if err := c.client.Bucket(initdata.Bucket).Put(initdata.Key, initdata.RawData, conType, s3.PublicReadWrite, s3.Options{}); err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		return err
	}

	return nil
}
