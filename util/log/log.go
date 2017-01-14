package log

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

//log level, from low to high, more high means more serious
const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	Ltime  = 1 << iota //time format "2006/01/02 15:04:05"
	Lfile              //file.go:123
	Llevel             //[Trace|Debug|Info...]
)

var LogLevelString = map[string]int{
	"trace": LevelTrace,
	"debug": LevelDebug,
	"info":  LevelInfo,
	"warn":  LevelWarn,
	"error": LevelError,
	"fatal": LevelFatal,
}

var LevelName [6]string = [6]string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

const TimeFormat = "2006/01/02 15:04:05"

const maxBufPoolSize = 16

const (
	NONE         = "\033[m"
	RED          = "\033[0;32;31m"
	LIGHT_RED    = "\033[1;31m"
	GREEN        = "\033[0;32;32m"
	LIGHT_GREEN  = "\033[1;32m"
	BLUE         = "\033[0;32;34m"
	LIGHT_BLUE   = "\033[1;34m"
	DARY_GRAY    = "\033[1;30m"
	CYAN         = "\033[0;36m"
	LIGHT_CYAN   = "\033[1;36m"
	PURPLE       = "\033[0;35m"
	LIGHT_PURPLE = "\033[1;35m"
	BROWN        = "\033[0;33m"
	YELLOW       = "\033[1;33m"
	LIGHT_GRAY   = "\033[0;37m"
	WHITE        = "\033[1;37m"
)

type Logger struct {
	sync.Mutex

	level int
	flag  int

	handler Handler

	quit chan struct{}
	msg  chan []byte

	bufs [][]byte

	wg sync.WaitGroup

	closed      bool
	fileHandler bool
}

//new a logger with specified handler and flag
func New(handler Handler, flag int) *Logger {
	var l = new(Logger)

	l.level = LevelInfo
	l.handler = handler

	l.flag = flag

	l.quit = make(chan struct{})
	l.closed = false
	l.fileHandler = false

	l.msg = make(chan []byte, 1024)

	l.bufs = make([][]byte, 0, 16)

	l.wg.Add(1)

	go l.run()

	return l
}

//new a default logger with specified handler and flag: Ltime|Lfile|Llevel
func NewDefault(handler Handler) *Logger {
	return New(handler, Ltime|Lfile|Llevel)
}

func newStdHandler() *StreamHandler {
	h, _ := NewStreamHandler(os.Stdout)
	return h
}

var std = NewDefault(newStdHandler())

func Close() {
	std.Close()
}

func (l *Logger) run() {
	defer l.wg.Done()
	for {
		select {
		case msg := <-l.msg:
			l.handler.Write(msg)
			l.putBuf(msg)
		case <-l.quit:
			if len(l.msg) == 0 {
				return
			}
		}
	}
}

func (l *Logger) popBuf() []byte {
	l.Lock()
	var buf []byte
	if len(l.bufs) == 0 {
		buf = make([]byte, 0, 1024)
	} else {
		buf = l.bufs[len(l.bufs)-1]
		l.bufs = l.bufs[0 : len(l.bufs)-1]
	}
	l.Unlock()

	return buf
}

func (l *Logger) putBuf(buf []byte) {
	l.Lock()
	if len(l.bufs) < maxBufPoolSize {
		buf = buf[0:0]
		l.bufs = append(l.bufs, buf)
	}
	l.Unlock()
}

func (l *Logger) Close() {
	if l.closed {
		return
	}
	l.closed = true

	close(l.quit)
	l.wg.Wait()
	l.quit = nil

	l.handler.Close()
}

//set log level, any log level less than it will not log
func (l *Logger) setLevel(level int) {
	l.level = level
}

func (l *Logger) Level() int {
	return l.level
}

func (l *Logger) SetLogFile(logFile string) {
	if logFile != "" {
		fileHandler, err := NewTimeRotatingFileHandler(logFile, WhenDay, 1)
		if err != nil {
			return
		}
		l.handler = fileHandler
		l.fileHandler = true
	}
}

//a low interface, maybe you can use it for your special log format
//but it may be not exported later......
func (l *Logger) Output(callDepth int, level int, format string, v ...interface{}) {
	if l.level > level {
		return
	}

	buf := l.popBuf()
	buf = append(buf, l.colorStart(level)...)

	if l.flag&Ltime > 0 {
		now := time.Now().Format(TimeFormat)
		buf = append(buf, now...)
		buf = append(buf, " - "...)
	}

	if l.flag&Llevel > 0 {
		buf = append(buf, LevelName[level]...)
		buf = append(buf, " - "...)
	}

	if l.flag&Lfile > 0 {
		pc, file, line, ok := runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		} else {
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					file = file[i+1:]
					break
				}
			}
		}

		buf = append(buf, file...)
		buf = append(buf, ":["...)

		buf = strconv.AppendInt(buf, int64(line), 10)
		buf = append(buf, "] - "...)

		if pc != 0 {
			f := runtime.FuncForPC(pc)
			if f != nil {
				funcNamePath := strings.Split(f.Name(), "/")
				buf = append(buf, "["+funcNamePath[len(funcNamePath)-1]+"] - "...)
			}
		}
	}

	s := fmt.Sprintf(format, v...)

	buf = append(buf, s...)

	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf = append(buf, '\n')
	}

	buf = append(buf, l.colorStop(level)...)

	l.msg <- buf
}

func (l *Logger) colorStart(level int) string {
	if l.fileHandler {
		return ""
	}

	if level >= LevelError {
		return RED
	} else if level == LevelWarn {
		return YELLOW
	}

	return ""
}
func (l *Logger) colorStop(level int) string {
	if l.fileHandler {
		return ""
	}

	if level >= LevelError {
		return NONE
	} else if level == LevelWarn {
		return NONE
	}

	return ""
}

func setLevel(level int) {
	std.setLevel(level)
}

func SetLogFile(logFile string) {
	std.SetLogFile(logFile)
}

func SetLevelS(level string) {
	SetLevel(LogLevelString[strings.ToLower(level)])
}

func StdLogger() *Logger {
	return std
}

func GetLevel() int {
	return std.level
}
