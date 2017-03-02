package swift

import (
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/result"
	"github.com/gin-gonic/gin"
)

func TestNew(t *testing.T) {
	type args struct {
		config *config.Config
	}
	tests := []struct {
		name string
		args args
		want *Swift
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doRequest(t *testing.T) {
	type args struct {
		method string
		url    string
		token  string
		body   io.Reader
		client *http.Client
	}
	tests := []struct {
		name     string
		args     args
		wantResp *http.Response
		wantErr  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := doRequest(tt.args.method, tt.args.url, tt.args.token, tt.args.body, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("doRequest() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestSwift_Auth(t *testing.T) {
	tests := []struct {
		name    string
		swt     *Swift
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.swt.Auth(); (err != nil) != tt.wantErr {
				t.Errorf("Swift.Auth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_buildUrl(t *testing.T) {
	type args struct {
		addr    string
		account string
		bucket  string
		key     string
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
			if got := buildUrl(tt.args.addr, tt.args.account, tt.args.bucket, tt.args.key); got != tt.want {
				t.Errorf("buildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildBucketUrl(t *testing.T) {
	type args struct {
		addr    string
		account string
		bucket  string
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
			if got := buildBucketUrl(tt.args.addr, tt.args.account, tt.args.bucket); got != tt.want {
				t.Errorf("buildBucketUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildAuthReq(t *testing.T) {
	type args struct {
		config *config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := buildAuthReq(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildAuthReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("buildAuthReq() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("buildAuthReq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSwift_request(t *testing.T) {
	type args struct {
		data   io.Reader
		method string
		url    string
		res    *result.Result
		ctx    *gin.Context
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
			tt.swt.request(tt.args.data, tt.args.method, tt.args.url, tt.args.res, tt.args.ctx)
		})
	}
}

func TestSwift_bucketrequest(t *testing.T) {
	type args struct {
		data   io.Reader
		method string
		url    string
		res    *result.Result
		ctx    *gin.Context
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
			tt.swt.bucketrequest(tt.args.data, tt.args.method, tt.args.url, tt.args.res, tt.args.ctx)
		})
	}
}
