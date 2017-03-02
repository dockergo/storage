package swift

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSwift_ListBuckets(t *testing.T) {
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
			tt.swt.ListBuckets(tt.args.ctx)
		})
	}
}
