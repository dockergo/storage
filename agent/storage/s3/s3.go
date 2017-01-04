package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/flyaways/storage/agent/config"
	"github.com/flyaways/storage/agent/storage/adapter"
)

type S3 struct {
	adapter.StorageAdapter
	config *config.Config
	client *s3.S3
}

func New(config *config.Config) *S3 {
	ns3 := new(S3)
	ns3.config = config
	ns3.Name = "s3"
	credentials := credentials.NewStaticCredentials(config.Storage.S3.AccessKey, config.Storage.S3.SecretKey, "")
	sess := session.New(&aws.Config{
		Region:           aws.String("cn-north-1"),
		Credentials:      credentials,
		Endpoint:         &config.Storage.S3.Addr,
		DisableSSL:       aws.Bool(true),
		LogLevel:         aws.LogLevel(0), //1 is open while 0 is close
		S3ForcePathStyle: aws.Bool(true),
	})
	ns3.client = s3.New(sess)
	return ns3
}
