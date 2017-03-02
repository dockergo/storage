package s3

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestS3c_ListBuckets(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		c    *S3c
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.ListBuckets(tt.args.ctx)
		})
	}
}
