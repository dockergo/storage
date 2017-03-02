package swift

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSwift_GetBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		swt  *Swift
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.swt.GetBucket(tt.args.ctx)
		})
	}
}

func TestSwift_PostBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		swt  *Swift
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.swt.PostBucket(tt.args.ctx)
		})
	}
}

func TestSwift_PutBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		swt  *Swift
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.swt.PutBucket(tt.args.ctx)
		})
	}
}

func TestSwift_HeadBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		swt  *Swift
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.swt.HeadBucket(tt.args.ctx)
		})
	}
}

func TestSwift_DeleteBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		swt  *Swift
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.swt.DeleteBucket(tt.args.ctx)
		})
	}
}

func TestSwift_bucket(t *testing.T) {
	type args struct {
		method string
		ctx    *gin.Context
	}
	tests := []struct {
		name string
		swt  *Swift
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.swt.bucket(tt.args.method, tt.args.ctx)
		})
	}
}
