package api_test

import "flag"

const (
	TimeFormat  = "Mon, 02 Jan 2006 15:04:05 GMT"
	Newfilename = "Newfilename"
	Common      = "\n*************************************************************************************************%s\n"
)

var accessKey = flag.String("accKey", "1GL02rRYQxK8s7FQh8dV", "accessKey")
var secretKey = flag.String("secKey", "2IDjaPOpFfkq5Zf9K4tKu8k5AKApY8S8eKV1zsRl", "secretKey")
var Types = flag.String("typ", "service", "service")
var Method = flag.String("method", "GET", "GET,POST,PUT,HEAD,DELETE,OPTIONS")
var addr = flag.String("addr", "127.0.0.1:20808", "addr")

var content = flag.String("ctx", "Stand on the shoulders of giants", "PUT method content")
var curfile = flag.String("curf", "curfile", "current file name")

var bucketName = flag.String("bktName", "bk5buckets", "bucket Name")
var key = flag.String("key", "keydefault", "key")
var newName = flag.String("newName", "newName", "newfile Name")
