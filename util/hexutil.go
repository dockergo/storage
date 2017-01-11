package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func GetMd5(raw []byte) []byte {
	m := md5.New()
	m.Write(raw)
	return m.Sum(nil)
}

func GetBase64Md5(raw []byte) string {
	return base64.StdEncoding.EncodeToString(GetMd5(raw))
}

func GetETagValue(raw []byte) string {
	return hex.EncodeToString(GetMd5(raw))
}

func GetETag(raw []byte) string {
	return fmt.Sprintf("\"%s\"", GetETagValue(raw))
}

func GetSha1(raw []byte) []byte {
	h := sha1.New()
	h.Write(raw)
	return h.Sum(nil)
}

func GetSha1Hex(raw []byte) string {
	return hex.EncodeToString(GetSha1(raw))
}

func MakeHmac(key []byte, data []byte) []byte {
	hash := hmac.New(sha1.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}
func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}
