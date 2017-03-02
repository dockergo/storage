package qiniu

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestQiniu_GetBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		c    *Qiniu
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.GetBucket(tt.args.ctx)
		})
	}
}
