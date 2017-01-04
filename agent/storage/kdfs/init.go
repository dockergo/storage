package kdfs

import (
	"bytes"

	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util/log"
)

func (kfs *Kdfs) InitBucket(initdata *adapter.InitData) error {
	url := buildBucketUrl(kfs.config.Storage.Kdfs.Addr, kfs.config.Storage.Kdfs.Account, initdata.Bucket)
	_, err := doRequest("PUT", url, nil, kfs.httpClient)
	if err != nil {
		log.Error("[%s:%s]", kfs.Name, err.Error())
		return err
	}

	return nil
}

func (kfs *Kdfs) InitObject(initdata *adapter.InitData) error {
	url := buildUrl(kfs.config.Storage.Kdfs.Addr, kfs.config.Storage.Kdfs.Account, initdata.Bucket, initdata.Key)
	_, err := doRequest("PUT", url, bytes.NewReader(initdata.RawData), kfs.httpClient)
	if err != nil {
		log.Error("[%s:%s]", kfs.Name, err.Error())
		return err
	}

	return nil
}
