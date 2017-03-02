package middleware

import (
	"reflect"
	"testing"

	"github.com/flyaways/storage/config"
	"github.com/gin-gonic/gin"
)

func TestAuthority(t *testing.T) {
	type args struct {
		credential *config.CredentialConfig
	}
	tests := []struct {
		name string
		args args
		want gin.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Authority(tt.args.credential); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignAuth_Auth(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		s    *SignAuth
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Auth(tt.args.ctx)
		})
	}
}

func TestSignAuth_sign(t *testing.T) {
	tests := []struct {
		name    string
		s       *SignAuth
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.sign(); (err != nil) != tt.wantErr {
				t.Errorf("SignAuth.sign() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignAuth_buildTime(t *testing.T) {
	tests := []struct {
		name    string
		s       *SignAuth
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.buildTime(); (err != nil) != tt.wantErr {
				t.Errorf("SignAuth.buildTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignAuth_buildCanonicalHeaders(t *testing.T) {
	tests := []struct {
		name string
		s    *SignAuth
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.buildCanonicalHeaders()
		})
	}
}

func TestSignAuth_buildCanonicalResource(t *testing.T) {
	tests := []struct {
		name string
		s    *SignAuth
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.buildCanonicalResource()
		})
	}
}

func TestSignAuth_buildStringToSign(t *testing.T) {
	tests := []struct {
		name    string
		s       *SignAuth
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.buildStringToSign(); (err != nil) != tt.wantErr {
				t.Errorf("SignAuth.buildStringToSign() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignAuth_buildSignature(t *testing.T) {
	tests := []struct {
		name string
		s    *SignAuth
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.buildSignature()
		})
	}
}

func TestSignAuth_getValues(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name  string
		s     *SignAuth
		args  args
		want  string
		want1 bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getValues(tt.args.keys...)
			if got != tt.want {
				t.Errorf("SignAuth.getValues() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SignAuth.getValues() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSignAuth_getOldSign(t *testing.T) {
	tests := []struct {
		name  string
		s     *SignAuth
		want  string
		want1 string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.getOldSign()
			if got != tt.want {
				t.Errorf("SignAuth.getOldSign() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SignAuth.getOldSign() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEscapePath(t *testing.T) {
	type args struct {
		path      string
		encodeSep bool
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
			if got := EscapePath(tt.args.path, tt.args.encodeSep); got != tt.want {
				t.Errorf("EscapePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
