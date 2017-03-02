package protocol

import (
	"reflect"
	"testing"

	"github.com/flyaways/storage/result"
	"github.com/gin-gonic/gin"
)

func TestHeader(t *testing.T) {
	type args struct {
		ctx    *gin.Context
		res    *result.Result
		bucket string
		key    string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Header(tt.args.ctx, tt.args.res, tt.args.bucket, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Header() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Header() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetHeader(t *testing.T) {
	type args struct {
		ctx     *gin.Context
		content []byte
		res     *result.Result
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetHeader(tt.args.ctx, tt.args.content, tt.args.res)
		})
	}
}

func TestKeyMaker(t *testing.T) {
	type args struct {
		ctx  *gin.Context
		key  string
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyMaker(tt.args.ctx, tt.args.key, tt.args.data); got != tt.want {
				t.Errorf("KeyMaker() = %v, want %v", got, tt.want)
			}
		})
	}
}
