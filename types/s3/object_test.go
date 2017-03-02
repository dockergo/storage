package s3

import (
	"testing"

	"github.com/flyaways/storage/result"
	"github.com/gin-gonic/gin"
)

func TestS3c_PutObject(t *testing.T) {
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
			tt.c.PutObject(tt.args.ctx)
		})
	}
}

func TestS3c_PostObject(t *testing.T) {
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
			tt.c.PostObject(tt.args.ctx)
		})
	}
}

func TestS3c_upload(t *testing.T) {
	type args struct {
		ctx      *gin.Context
		res      *result.Result
		data     []byte
		finalkey string
		bucket   string
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
			tt.c.upload(tt.args.ctx, tt.args.res, tt.args.data, tt.args.finalkey, tt.args.bucket)
		})
	}
}

func TestS3c_GetObject(t *testing.T) {
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
			tt.c.GetObject(tt.args.ctx)
		})
	}
}

func TestS3c_HeadObject(t *testing.T) {
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
			tt.c.HeadObject(tt.args.ctx)
		})
	}
}

func TestS3c_DeleteObject(t *testing.T) {
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
			tt.c.DeleteObject(tt.args.ctx)
		})
	}
}
