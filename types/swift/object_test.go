package swift

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSwift_PutObject(t *testing.T) {
	type args struct {
		ctx *gin.Context
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
			tt.swt.PutObject(tt.args.ctx)
		})
	}
}

func TestSwift_PostObject(t *testing.T) {
	type args struct {
		ctx *gin.Context
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
			tt.swt.PostObject(tt.args.ctx)
		})
	}
}

func TestSwift_HeadObject(t *testing.T) {
	type args struct {
		ctx *gin.Context
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
			tt.swt.HeadObject(tt.args.ctx)
		})
	}
}

func TestSwift_GetObject(t *testing.T) {
	type args struct {
		ctx *gin.Context
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
			tt.swt.GetObject(tt.args.ctx)
		})
	}
}

func TestSwift_DeleteObject(t *testing.T) {
	type args struct {
		ctx *gin.Context
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
			tt.swt.DeleteObject(tt.args.ctx)
		})
	}
}

func TestSwift_object(t *testing.T) {
	type args struct {
		method string
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
			tt.swt.object(tt.args.method, tt.args.ctx)
		})
	}
}
