package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code       string
	Message    string
	StatusCode int
}

func NewError(code, message string, httpCode int) *Error {
	return &Error{code, message, httpCode}
}

func (e *Error) Error() string {
	return fmt.Sprintf("<Code>%s</Code>\n<Message>%s</Message>", e.Code, e.Message)
}

var (
	MissingContentLength  = NewError("MissingContentLength", "You must provide the Content-Length HTTP header.", http.StatusLengthRequired)
	SignatureDoesNotMatch = NewError("SignatureDoesNotMatch", "The request signature we calculated does not match the signature ", http.StatusForbidden)
	NoSuchKey             = NewError("NoSuchKey", "The specified key does not exist.", http.StatusNotFound)
	InvalidRange          = NewError("InvalidRange", "The requested range cannot be satisfied.", http.StatusRequestedRangeNotSatisfiable)
	InternalError         = NewError("InternalError", "We encountered an internal error. Please try again.", http.StatusInternalServerError)
	URLExpired            = NewError("URLExpired", "The URL is expired, please get a new one.", http.StatusForbidden)
	AccessDenied          = NewError("AccessDenied", "Access Denied", http.StatusForbidden)
	InvalidArgument       = NewError("InvalidArgument", "Invalid Argument", http.StatusBadRequest)
	InvalidKey            = NewError("invalidkey", "invalid key", http.StatusBadRequest)
	InvalidAccessKeyId    = NewError("InvalidAccessKeyId", "The KS3 access key Id you provided does not exist in our records.", http.StatusForbidden)

	UnknownError    = NewError("UnknownError", "UnknownError", http.StatusInternalServerError)
	BucketError     = NewError("BucketError", "BucketError", http.StatusBadRequest)
	MD5DoesNotMatch = NewError("MD5DoesNotMatch", "MD5DoesNotMatch", http.StatusBadRequest)
	FileEmpty       = NewError("FileEmpty", "FileEmpty", http.StatusBadRequest)
	ReqBodyNil      = NewError("ReqBodyNil", "request body nil", http.StatusBadRequest)

	AuthorizationNull  = NewError("AuthorizationNull", "require Authorization header", http.StatusUnauthorized)
	ErrCSRFRequire     = NewError("CSRFRequire", "csrfRequire", http.StatusForbidden)
	ErrApiAccessDenied = NewError("AccessDenied", "ApiAccessDenied", http.StatusForbidden)
	ErrInvalidArgument = NewError("InvalidArgument", "InvalidArgument", http.StatusBadRequest)

	NoSuchUpload     = NewError("NoSuchUpload", "指定的分块上传任务不存在。可能是上传ID无效，也可能是分块上传任务已经完成或放弃。", http.StatusNotFound)
	InvalidPartOrder = NewError("InvalidPartOrder", "分块列表没有按照升序排列。分块必须按序指定分块序号。", http.StatusBadRequest)
	InvalidPart      = NewError("InvalidPartOrder", "一个或多个指定的块没有找到。块可能没有被上传，也可能因为上传了但实体标签不匹配。", http.StatusBadRequest)
	EntityTooSmall   = NewError("InvalidPartOrder", "用户拟上传的块大小小于对象所允许的最小值。除了最后一个块之外，每一个块至少在5MB以上。当文件总大小在5M以内的话，可以允许除了最后一个块之外，每一个块至少在100K以上。", http.StatusBadRequest)

	NoSuchBucket     = NewError("NoSuchBucket", "The specified bucket does not exist.", http.StatusBadRequest)
	BucketExist      = NewError("BucketExist", "The specified bucket is exist.", http.StatusConflict)
	BucketNotEmpty   = NewError("BucketNotEmpty", "The specified bucket does not empty.", http.StatusConflict)
	UnSupportError   = NewError("UnSupportError", "UnSupportError", http.StatusInternalServerError)
	PermissionDenied = NewError("PermissionDenied", "permission denied of write and delete.", http.StatusInternalServerError)

	StatusAccepted  = NewError("StatusAccepted", "StatusAccepted", http.StatusAccepted)
	StatusNotFound  = NewError("StatusNotFound", "StatusNotFound", http.StatusNotFound)
	StatusNoContent = NewError("StatusNoContent", "StatusNoContent", http.StatusNoContent)
	StatusConflict  = NewError("StatusConflict", "StatusConflict", http.StatusConflict)
)
