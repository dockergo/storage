package s3

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/flyaways/storage/agent/constant"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/result"
	"github.com/flyaways/storage/agent/util/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func (s *S3) PutObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	rawRequestdata, finalkey, err := protocol.PutHeadchecker(ctx, res, bucket, key)
	if err != nil {
		return
	}

	s.uploadObject(ctx, res, rawRequestdata, finalkey, bucket)
}

func (s *S3) PostObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParamPost(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	rawRequestdata, finalkey, err := protocol.PostHeadchecker(ctx, res, bucket, key)
	if err != nil {
		return
	}

	s.uploadObject(ctx, res, rawRequestdata, finalkey, bucket)
}

func (s *S3) uploadObject(ctx *gin.Context, res *result.Result, rawRequestdata []byte, finalkey, bucket string) {
	contentType := ctx.Request.Header.Get(constant.ContentType)
	params := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(finalkey),
		Body:        bytes.NewReader(rawRequestdata),
		ContentType: aws.String(contentType),
	}

	_, s3Err := s.client.PutObject(params)
	if s3Err != nil {
		log.Error("[%s:%s]", s.Name, s3Err.Error())
		res.Error(s3Err)
	}
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: finalkey})
	return
}

func (s *S3) GetObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	s3Resp, s3Err := s.client.GetObject(params)
	if s3Err != nil {
		log.Error("[%s:%s]", s.Name, s3Err.Error())
		res.Error(s3Err)
	}
	content, _ := ioutil.ReadAll(s3Resp.Body)
	protocol.GetCkecker(ctx, content, res)
}

func (s *S3) HeadObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}
	params := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	s3Resp, s3Err := s.client.HeadObject(params)
	if s3Err != nil {
		log.Error("[%s:%s]", s.Name, s3Err.Error())
		res.Error(s3Err)
	}

	objEtag := s3Resp.ETag
	if protocol.CheckETag(ctx, *objEtag) {
		log.Error("[%s:S3 HeadObject invalid Etag]]", s.Name)
		res.Error("[S3 HeadObject invalid Etag]")
		return
	}

	protocol.HeadChecker(ctx, res, *objEtag, (*s3Resp.LastModified).Format(constant.TimeFormat))
	ctx.Status(http.StatusOK)
}

func (b *S3) DeleteObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	_, s3Err := b.client.DeleteObject(params)
	if s3Err != nil {
		log.Error("[s3 DeleteObject error:%s]", s3Err.Error())
		res.Error(s3Err)
		return
	}

	ctx.Status(http.StatusOK)
}
