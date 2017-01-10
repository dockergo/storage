package api_test

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/flyaways/tracker"
)

func DoSignature(HTTPVerb, ContentMD5, ContentType, Date, CanonicalizedResource, secretKey string,
	CanonicalizedKssHeaders map[string]string) string {
	stringToSign := HTTPVerb + "\n" +
		ContentMD5 + "\n" +
		ContentType + "\n" +
		Date + "\n"

	for k, v := range CanonicalizedKssHeaders {
		stringToSign += k + ":" + v + "\n"
	}

	stringToSign += CanonicalizedResource
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(stringToSign))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))

	fmt.Printf("[%40s:\t%-50s]\n", tracker.Blue("signature"), tracker.Yellow(string(sign)))

	return sign
}

type Policy struct {
	Expiration string
	Conditions map[string]string
}

func (p *Policy) Marshal() []byte {
	json := `{"expiration": "` + p.Expiration + `","conditions": [`
	for k, v := range p.Conditions {
		json += `{"` + k + `": "` + v + `"},`
	}
	json += `]}`
	return []byte(json)
}

func GetExpireTime() string {
	expires := time.Now().Unix() + 600
	return strconv.FormatInt(expires, 10)
}

func GetDate() string {
	return time.Now().UTC().Format(TimeFormat)
}

func MakeHmac(key []byte, data []byte) []byte {
	hash := hmac.New(sha1.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}
