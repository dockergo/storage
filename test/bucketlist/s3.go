package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

var bkt = flag.String("bucket", "wpsfiletest", "bucket")

func main() {
	flag.Parse()
	auth, err := aws.GetAuth("CUPQ9OHNSEQ6KYDLYNK2", "jmgkfRgE2IkgKH9gmOoboOanID7SNHQ2yNRHw6GF")
	if err != nil {
		log.Fatal(err)
	}

	var cnc = aws.Region{
		S3Endpoint:           "http://192.168.20.4:8888",
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
	}

	client := s3.New(auth, cnc)
	resp, err := client.ListBuckets()
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range resp.Buckets {
		fmt.Printf("[Bucket: %s]\n", bucket.Name)
		if *bkt == bucket.Name {
			keys, err := bucket.GetBucketContents()
			if err == nil {
				for _, key := range *keys {
					fmt.Printf("%s\t%s\t%10.d\t%s\t%s\t%s\t%s\n",
						key.Key, key.LastModified, key.Size, key.ETag,
						key.StorageClass, key.Owner.ID, key.Owner.DisplayName)
				}
			}
		}

	}
}
