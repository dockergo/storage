package oss

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestOSS_ListBuckets(t *testing.T) {
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
			tt.ossc.ListBuckets(tt.args.ctx)
		})
	}
}
