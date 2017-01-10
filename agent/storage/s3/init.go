package s3

import (
	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util/log"
	"github.com/mitchellh/goamz/s3"
)

func (c *s3c) InitBucket(initdata *adapter.InitData) error {
	if err := c.client.Bucket(initdata.Bucket).PutBucket(s3.PublicRead); err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		return err
	}

	return nil
}

func (c *s3c) InitObject(initdata *adapter.InitData) error {
	conType := "application/ocet-stream"
	if err := c.client.Bucket(initdata.Bucket).Put(initdata.Key, initdata.RawData, conType, s3.PublicReadWrite); err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		return err
	}

	return nil
}
