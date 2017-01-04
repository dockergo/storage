package s3

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
)

func (s *S3) PutBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}
	_, s3Err := s.client.CreateBucket(params)
	if s3Err != nil {
		log.Error("[%s:%s]", s.Name, s3Err.Error())
		res.Error(s3Err)
	}
	ctx.Status(http.StatusOK)

}

func (s *S3) HeadBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	params := &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	}
	_, s3Err := s.client.HeadBucket(params)
	if s3Err != nil {
		log.Error("[%s:%s]", s.Name, s3Err.Error())
		res.Error(s3Err)
	}
	ctx.Status(http.StatusOK)

}

func (s *S3) DeleteBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	params := &s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	}
	_, s3Err := s.client.DeleteBucket(params)
	if s3Err != nil {
		log.Error("[%s:%s]", s.Name, s3Err.Error())
		res.Error(s3Err)
	}
	ctx.Status(http.StatusOK)

}
