package qiniu

import (
	"reflect"
	"testing"

	"github.com/flyaways/storage/config"
)

func TestNew(t *testing.T) {
	type args struct {
		config *config.Config
	}
	tests := []struct {
		name string
		args args
		want *Qiniu
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
