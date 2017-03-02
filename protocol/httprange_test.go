package protocol

import (
	"net/textproto"
	"reflect"
	"testing"
)

func TestHttpRange_ContentRange(t *testing.T) {
	type args struct {
		size int64
	}
	tests := []struct {
		name string
		r    HttpRange
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.ContentRange(tt.args.size); got != tt.want {
				t.Errorf("HttpRange.ContentRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpRange_MimeHeader(t *testing.T) {
	type args struct {
		contentType string
		size        int64
	}
	tests := []struct {
		name string
		r    HttpRange
		args args
		want textproto.MIMEHeader
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.MimeHeader(tt.args.contentType, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpRange.MimeHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRange(t *testing.T) {
	type args struct {
		s    string
		size int64
	}
	tests := []struct {
		name    string
		args    args
		want    []HttpRange
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRange(tt.args.s, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumRangesSize(t *testing.T) {
	type args struct {
		ranges []HttpRange
	}
	tests := []struct {
		name     string
		args     args
		wantSize int64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSize := SumRangesSize(tt.args.ranges); gotSize != tt.wantSize {
				t.Errorf("SumRangesSize() = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}
