package posix

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPosix_GetBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name  string
		posix *Posix
		args  args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.posix.GetBucket(tt.args.ctx)
		})
	}
}

func TestPosix_PutBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name  string
		posix *Posix
		args  args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.posix.PutBucket(tt.args.ctx)
		})
	}
}

func TestPosix_HeadBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name  string
		posix *Posix
		args  args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.posix.HeadBucket(tt.args.ctx)
		})
	}
}

func TestPosix_DeleteBucket(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name  string
		posix *Posix
		args  args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.posix.DeleteBucket(tt.args.ctx)
		})
	}
}
