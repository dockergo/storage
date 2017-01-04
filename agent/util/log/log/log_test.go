package log

import (
	"os"
	"testing"

	"github.com/getsentry/raven-go"
)

func TestStdStreamLog(t *testing.T) {
	h, _ := NewStreamHandler(os.Stdout)
	s := NewDefault(h)
	s.Info("hello world")

	s.Close()

	Info("hello world")
}

func TestRotatingFileLog(t *testing.T) {
	path := "./test_log"
	os.RemoveAll(path)

	os.Mkdir(path, 0777)
	fileName := path + "/test"

	h, err := NewRotatingFileHandler(fileName, 10, 2)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 10)

	h.Write(buf)

	h.Write(buf)

	if _, err := os.Stat(fileName + ".1"); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(fileName + ".2"); err == nil {
		t.Fatal(err)
	}

	h.Write(buf)
	if _, err := os.Stat(fileName + ".2"); err != nil {
		t.Fatal(err)
	}

	h.Close()

	os.RemoveAll(path)
}

func TestSentry(t *testing.T) {
	l := NewDefault(newStdHandler())
	c, err := raven.NewClient("http://226d9ac62029427180aea0d56414c813:e931de0aef5f49df9009c0c8b1d21a1f@10.20.132.153:9000/2", nil)
	if err != nil {
		t.Error("raven.NewClient fail:" + err.Error())
	}
	l.SetSentry(c)
	l.Warn("test warnning")
	l.Error("test error")
	l.Fatal("test fatal")
}
