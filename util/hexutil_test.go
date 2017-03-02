package util

import (
	"reflect"
	"testing"
)

func TestGetMd5(t *testing.T) {
	type args struct {
		raw []byte
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
			if got := GetMd5(tt.args.raw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBase64Md5(t *testing.T) {
	type args struct {
		raw []byte
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
			if got := GetBase64Md5(tt.args.raw); got != tt.want {
				t.Errorf("GetBase64Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetETagValue(t *testing.T) {
	type args struct {
		raw []byte
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
			if got := GetETagValue(tt.args.raw); got != tt.want {
				t.Errorf("GetETagValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetETag(t *testing.T) {
	type args struct {
		raw []byte
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
			if got := GetETag(tt.args.raw); got != tt.want {
				t.Errorf("GetETag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSha1(t *testing.T) {
	type args struct {
		raw []byte
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
			if got := GetSha1(tt.args.raw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSha1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSha1Hex(t *testing.T) {
	type args struct {
		raw []byte
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
			if got := GetSha1Hex(tt.args.raw); got != tt.want {
				t.Errorf("GetSha1Hex() = %v, want %v", got, tt.want)
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
