package storage

import (
	"reflect"
	"testing"

	"github.com/flyaways/storage/app"
	"github.com/gin-gonic/gin"
)

func TestPolicy_Marshal(t *testing.T) {
	type fields struct {
		Expiration string
		Conditions map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plc := &Policy{
				Expiration: tt.fields.Expiration,
				Conditions: tt.fields.Conditions,
			}
			if got := plc.Marshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Policy.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetExpireTime(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetExpireTime(); got != tt.want {
				t.Errorf("GetExpireTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDate(); got != tt.want {
				t.Errorf("GetDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeHmac(t *testing.T) {
	type args struct {
		key  []byte
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeHmac(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeHmac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64Encode(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Encode(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Base64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walkDir(t *testing.T) {
	type args struct {
		dirPth string
		suffix string
	}
	tests := []struct {
		name      string
		args      args
		wantFiles []string
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, err := walkDir(tt.args.dirPth, tt.args.suffix)
			if (err != nil) != tt.wantErr {
				t.Errorf("walkDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("walkDir() = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

func TestDoSignature(t *testing.T) {
	type args struct {
		HTTPVerb                string
		ContentMD5              string
		ContentType             string
		Date                    string
		CanonicalizedResource   string
		secretKey               string
		CanonicalizedKssHeaders map[string]string
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
			if got := DoSignature(tt.args.HTTPVerb, tt.args.ContentMD5, tt.args.ContentType, tt.args.Date, tt.args.CanonicalizedResource, tt.args.secretKey, tt.args.CanonicalizedKssHeaders); got != tt.want {
				t.Errorf("DoSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_objectPost(t *testing.T) {
	type args struct {
		app        *app.App
		router     *gin.Engine
		curfile    string
		bucketName string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			objectPost(tt.args.app, tt.args.router, tt.args.curfile, tt.args.bucketName)
		})
	}
}

func Test_initObject(t *testing.T) {
	type args struct {
		app    *app.App
		router *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initObject(tt.args.app, tt.args.router)
		})
	}
}

func Test_initBucket(t *testing.T) {
	type args struct {
		app    *app.App
		router *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initBucket(tt.args.app, tt.args.router)
		})
	}
}
