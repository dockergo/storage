package oss

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestOSS_PutBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		ossc *OSS
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ossc.PutBucket(tt.args.ctx)
		})
	}
}

func TestOSS_PostBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		ossc *OSS
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ossc.PostBucket(tt.args.ctx)
		})
	}
}

func TestOSS_DeleteBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		ossc *OSS
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ossc.DeleteBucket(tt.args.ctx)
		})
	}
}

func TestOSS_HeadBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		ossc *OSS
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ossc.HeadBucket(tt.args.ctx)
		})
	}
}

func TestOSS_GetBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		ossc *OSS
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ossc.GetBucket(tt.args.ctx)
		})
	}
}
