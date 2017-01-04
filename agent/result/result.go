package result

import (
	"bytes"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/flyaways/storage/agent/errors"
)

type Result struct {
	context *gin.Context
}

func NewResult(context *gin.Context) *Result {
	return &Result{context: context}
}

func (r *Result) Json(code int, data interface{}) {
	r.context.JSON(code, data)
}

func (r *Result) Xml(code int, data interface{}) {
	r.context.XML(code, data)
}

func (r *Result) Data(code int, contentType string, data []byte) {
	r.context.Data(code, contentType, data)
}

const (
	XMLHeader = `<?xml version="1.0" encoding="utf-8" standalone="yes"?>` + "\n"
)

type ErrorResult struct {
	XMLName   xml.Name `xml:"Error"`
	Code      string   `xml:"Code"`
	Message   string   `xml:"Message"`
	Resource  string   `xml:"Resource"`
	RequestId string   `xml:"RequestId"`
}

func (r *Result) Error(errorData interface{}) {

	requestId := r.context.MustGet("requestId")

	errResult := ErrorResult{Code: errors.UnknownError.Code,
		Message:   errors.UnknownError.Message,
		Resource:  r.context.Request.URL.Path,
		RequestId: requestId.(string)}

	statusCode := 500

	switch e := errorData.(type) {
	case *errors.Error:
		statusCode = e.StatusCode
		errResult = ErrorResult{Code: e.Code,
			Message:   e.Message,
			Resource:  r.context.Request.URL.Path,
			RequestId: requestId.(string)}

	case error: // unknown error
		statusCode = errors.UnknownError.StatusCode
		errResult = ErrorResult{Code: errors.UnknownError.Code,
			Message:   e.Error(),
			Resource:  r.context.Request.URL.Path,
			RequestId: requestId.(string)}

	}

	var buf bytes.Buffer
	rs := []byte(XMLHeader)
	data, err := xml.MarshalIndent(errResult, "  ", "    ")
	if err != nil {
		panic("xml.MarshalIndent error")
	}
	buf.Write(rs)
	buf.Write(data)
	r.context.Data(statusCode, "application/xml;charset=utf-8", buf.Bytes())
}
