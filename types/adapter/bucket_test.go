package adapter

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStorageAdapter_GetBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		s    *StorageAdapter
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.GetBucket(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_PutBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		s    *StorageAdapter
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.PutBucket(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_HeadBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		s    *StorageAdapter
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.HeadBucket(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_DeleteBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		s    *StorageAdapter
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.DeleteBucket(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_OptionsBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		s    *StorageAdapter
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.OptionsBucket(tt.args.ctx)
		})
	}
}
