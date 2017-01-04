package s3

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util/log"
)

func (s *S3) InitBucket(initdata *adapter.InitData) error {
	params := &s3.CreateBucketInput{
		Bucket: aws.String(initdata.Bucket),
	}

	if _, err := s.client.CreateBucket(params); err != nil {
		log.Error("[%s:%s]", s.Name, err.Error())
		return err
	}

	return nil
}

func (s *S3) InitObject(initdata *adapter.InitData) error {
	conType := "application/ocet-stream"
	params := &s3.PutObjectInput{
		Bucket:      aws.String(initdata.Bucket),
		Key:         aws.String(initdata.Key),
		Body:        bytes.NewReader(initdata.RawData),
		ContentType: aws.String(conType),
	}

	if _, err := s.client.PutObject(params); err != nil {
		log.Error("[%s:%s]", s.Name, err.Error())
		return err
	}

	return nil
}
