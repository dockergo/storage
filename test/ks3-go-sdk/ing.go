package main

import (
	"fmt"
	"strings"

	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/aws/credentials"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
)

var bucket = string("aa-go-sdk")
var key = string("aws/config.go")
var endpoint = string("kss.ksyun.com")
var accesskey = string("lMQTr0hNlMpB0iOk/i+x")
var secretkey = string("D4CsYLs75JcWEjbiI22zR3P7kJ/+5B1qdEje7A7I")
var region = string("HANGZHOU")

/***
var bucket = string("wpsfile")
var key = string("b04f3ee8f5e43fa3b162981b50bb72fe1acabb33")
var endpoint = string("192.168.20.131:20808")
var accesskey = string("1GL02rRYQxK8s7FQh8dV")
var secretkey = string("2IDjaPOpFfkq5Zf9K4tKu8k5AKApY8S8eKV1zsRl")
var region = string("cn-north-1")
***/

func New(credentials *credentials.Credentials) *s3.S3 {
	svc := s3.New(&aws.Config{
		Region:           region,
		Credentials:      credentials,
		Endpoint:         endpoint,
		DisableSSL:       true,
		LogLevel:         1,
		S3ForcePathStyle: true,
		LogHTTPBody:      true,
	})
	return svc
}

func main() {
	credentials := credentials.NewStaticCredentials(accesskey, secretkey, "")
	svc := New(credentials)

	putObject(svc)
	fmt.Printf("\n\n\n")
	getObject(svc)
	fmt.Printf("\n\n\n")
	headObject(svc)
}

func putObject(c *s3.S3) {

	contenttype := "application/ocet-stream"
	out, err := c.PutObject(
		&s3.PutObjectInput{
			Bucket:      &bucket,
			Key:         &key,
			Body:        strings.NewReader("this is the test from wps"),
			ContentType: &contenttype,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func getObject(c *s3.S3) {
	out, err := c.GetObject(
		&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out.Metadata)
	fmt.Println(*out.ContentLength)
	fmt.Println(*out.ContentType)

	b := make([]byte, 20)
	n, err := out.Body.Read(b)
	fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)

}

func headObject(c *s3.S3) {
	out, err := c.HeadObject(
		&s3.HeadObjectInput{
			Bucket: &bucket,
			Key:    &key,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out.Metadata)
	fmt.Println(*out.ContentLength)
	fmt.Println(*out.ContentType)
}
