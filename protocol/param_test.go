package protocol

import (
	"reflect"
	"testing"

	"github.com/flyaways/storage/result"
	"github.com/gin-gonic/gin"
)

func TestParam(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name  string
		args  args
		want  *result.Result
		want1 string
		want2 string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Param(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Param() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Param() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Param() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
