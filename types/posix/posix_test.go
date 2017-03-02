package posix

import (
	"os"
	"reflect"
	"testing"

	"github.com/flyaways/storage/config"
	"github.com/flyaways/storage/errors"
)

func TestNew(t *testing.T) {
	type args struct {
		config *config.Config
	}
	tests := []struct {
		name string
		args args
		want *Posix
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

func TestPosix_IsExist(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name  string
		posix *Posix
		args  args
		want  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.posix.IsExist(tt.args.filename); got != tt.want {
				t.Errorf("Posix.IsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosix_DirChecker(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name  string
		posix *Posix
		args  args
		want  *errors.Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.posix.DirChecker(tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Posix.DirChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosix_FileChecker(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name  string
		posix *Posix
		args  args
		want  *errors.Error
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.posix.FileChecker(tt.args.filepath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Posix.FileChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosix_getBucketPath(t *testing.T) {
	type args struct {
		bucket string
	}
	tests := []struct {
		name  string
		posix *Posix
		args  args
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.posix.getBucketPath(tt.args.bucket); got != tt.want {
				t.Errorf("Posix.getBucketPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_walkDir(t *testing.T) {
	type args struct {
		dirPth string
	}
	tests := []struct {
		name          string
		args          args
		wantFiles     []os.FileInfo
		wantFilenames []string
		wantErr       bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFiles, gotFilenames, err := walkDir(tt.args.dirPth)
			if (err != nil) != tt.wantErr {
				t.Errorf("walkDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("walkDir() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
			}
			if !reflect.DeepEqual(gotFilenames, tt.wantFilenames) {
				t.Errorf("walkDir() gotFilenames = %v, want %v", gotFilenames, tt.wantFilenames)
			}
		})
	}
}
