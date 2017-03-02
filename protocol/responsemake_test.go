package protocol

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_responseMake(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			responseMake(tt.args.ctx)
		})
	}
}
