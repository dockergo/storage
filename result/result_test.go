package result

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewResult(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name string
		args args
		want *Result
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResult(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_Json(t *testing.T) {
	type args struct {
		code int
		data interface{}
	}
	tests := []struct {
		name string
		r    *Result
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Json(tt.args.code, tt.args.data)
		})
	}
}

func TestResult_Xml(t *testing.T) {
	type args struct {
		code int
		data interface{}
	}
	tests := []struct {
		name string
		r    *Result
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Xml(tt.args.code, tt.args.data)
		})
	}
}

func TestResult_Data(t *testing.T) {
	type args struct {
		code        int
		contentType string
		data        []byte
	}
	tests := []struct {
		name string
		r    *Result
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Data(tt.args.code, tt.args.contentType, tt.args.data)
		})
	}
}

func TestResult_Error(t *testing.T) {
	type args struct {
		errorData interface{}
	}
	tests := []struct {
		name string
		r    *Result
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Error(tt.args.errorData)
		})
	}
}
