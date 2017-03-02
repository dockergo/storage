package types

import (
	"reflect"
	"testing"

	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/types/adapter"
)

func TestNew(t *testing.T) {
	type args struct {
		config *config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    adapter.Storager
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
