package adapter

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStorageAdapter_ListBuckets(t *testing.T) {
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
			tt.s.ListBuckets(tt.args.ctx)
		})
	}
}
