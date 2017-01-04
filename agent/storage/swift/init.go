package swift

import (
	"bytes"

	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util/log"
)

func (swt *Swift) InitBucket(initdata *adapter.InitData) error {
	url := buildBucketUrl(swt.config.Storage.Swift.Addr, swt.authAccount, initdata.Bucket)
	_, err := doRequest("PUT", url, swt.authToken, nil, swt.httpClient)
	if err != nil {
		log.Error("[%s:%s]", swt.Name, err.Error())
		return err
	}

	return nil
}

func (swt *Swift) InitObject(initdata *adapter.InitData) error {
	url := buildUrl(swt.config.Storage.Swift.Addr, swt.authAccount, initdata.Bucket, initdata.Key)
	_, err := doRequest("PUT", url, swt.authToken, bytes.NewReader(initdata.RawData), swt.httpClient)
	if err != nil {
		log.Error("[%s:%s]", swt.Name, err.Error())
		return err
	}

	return nil
}
