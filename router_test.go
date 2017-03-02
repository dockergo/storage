package storage

import (
	"testing"

	"github.com/flyaways/storage/app"
	"github.com/gin-gonic/gin"
)

func Test_regRouters(t *testing.T) {
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
			regRouters(tt.args.app, tt.args.router)
		})
	}
}

func Test_greenGroup(t *testing.T) {
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
			greenGroup(tt.args.app, tt.args.router)
		})
	}
}

func Test_bucketGroup(t *testing.T) {
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
			bucketGroup(tt.args.app, tt.args.router)
		})
	}
}

func Test_objectGroup(t *testing.T) {
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
			objectGroup(tt.args.app, tt.args.router)
		})
	}
}

func Test_serviceGroup(t *testing.T) {
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
			serviceGroup(tt.args.app, tt.args.router)
		})
	}
}
