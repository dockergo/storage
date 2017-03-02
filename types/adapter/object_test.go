package adapter

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStorageAdapter_PutObject(t *testing.T) {
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
			tt.s.PutObject(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_PostObject(t *testing.T) {
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
			tt.s.PostObject(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_GetObject(t *testing.T) {
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
			tt.s.GetObject(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_HeadObject(t *testing.T) {
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
			tt.s.HeadObject(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_DeleteObject(t *testing.T) {
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
			tt.s.DeleteObject(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_MoveObject(t *testing.T) {
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
			tt.s.MoveObject(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_CopyObject(t *testing.T) {
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
			tt.s.CopyObject(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_FetchObject(t *testing.T) {
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
			tt.s.FetchObject(tt.args.ctx)
		})
	}
}

func TestStorageAdapter_OptionsObject(t *testing.T) {
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
			tt.s.OptionsObject(tt.args.ctx)
		})
	}
}
