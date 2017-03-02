package posix

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_walkDirs(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name        string
		args        args
		wantBuckets []string
		wantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBuckets, err := walkDirs(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("walkDirs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBuckets, tt.wantBuckets) {
				t.Errorf("walkDirs() = %v, want %v", gotBuckets, tt.wantBuckets)
			}
		})
	}
}

func TestPosix_ListBuckets(t *testing.T) {
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
			tt.posix.ListBuckets(tt.args.ctx)
		})
	}
}
