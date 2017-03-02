package util

import "testing"

func TestDetails(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			Details()
		})
	}
}

func TestExec(t *testing.T) {
	type args struct {
		name string
		arg  []string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Exec(tt.args.name, tt.args.arg...)
		})
	}
}
