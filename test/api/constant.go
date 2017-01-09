package api_test

import "flag"

const (
	TimeFormat  = "Mon, 02 Jan 2006 15:04:05 GMT"
	Newfilename = "Newfilename"
	Common      = "\n*************************************************************************************************%s\n"
)

var accessKey = flag.String("accessKey", "1GL02rRYQxK8s7FQh8dV", "accessKey")
var secretKey = flag.String("secretKey", "2IDjaPOpFfkq5Zf9K4tKu8k5AKApY8S8eKV1zsRl", "secretKey")
var Types = flag.String("type", "service", "service")
var Method = flag.String("method", "GET", "GET,POST,PUT,HEAD,DELETE,OPTIONS")
var addr = flag.String("addr", "127.0.0.1:20808", "addr")

var content = flag.String("content", "Stand on the shoulders of giants", "content")
var curfile = flag.String("curfile", "defaultcurrent", "currentfile name")

var bucketName = flag.String("bucketName", "defaultbucket", "bucketName")
var key = flag.String("key", "defaultkey", "key")
var newName = flag.String("newName", "defaultkey", "newName")
