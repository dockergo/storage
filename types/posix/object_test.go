package posix

import (
	"testing"

	"github.com/flyaways/storage/result"
	"github.com/gin-gonic/gin"
)

func TestPosix_PutObject(t *testing.T) {
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
			tt.posix.PutObject(tt.args.ctx)
		})
	}
}

func TestPosix_PostObject(t *testing.T) {
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
			tt.posix.PostObject(tt.args.ctx)
		})
	}
}

func TestPosix_uploadObject(t *testing.T) {
	type args struct {
		ctx      *gin.Context
		res      *result.Result
		data     []byte
		finalkey string
		bucket   string
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
			tt.posix.uploadObject(tt.args.ctx, tt.args.res, tt.args.data, tt.args.finalkey, tt.args.bucket)
		})
	}
}

func TestPosix_GetObject(t *testing.T) {
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
			tt.posix.GetObject(tt.args.ctx)
		})
	}
}

func TestPosix_HeadObject(t *testing.T) {
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
			tt.posix.HeadObject(tt.args.ctx)
		})
	}
}

func TestPosix_DeleteObject(t *testing.T) {
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
			tt.posix.DeleteObject(tt.args.ctx)
		})
	}
}
