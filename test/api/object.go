package api_test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/flyaways/tracker"
)

func ObjectPost() {
	urlStr := fmt.Sprintf("http://%s/%s", *addr, *bucketName)
	fmt.Printf("\n[%40s:\t%-50s]\n", tracker.Blue("OBJECT-%s-URL", "POST"), tracker.Yellow(urlStr))

	policy := &Policy{
		Expiration: time.Unix(time.Now().Add(time.Minute*30).Unix(), 0).UTC().Format("2006-01-02T15:04:05.000Z"),
		Conditions: map[string]string{
			"bucket": *bucketName,
			"key":    *curfile,
		}}

	json := policy.Marshal()
	signature := string(Base64Encode(MakeHmac([]byte(*secretKey), Base64Encode(json))))

	var buffer bytes.Buffer
	w := multipart.NewWriter(&buffer)
	w.WriteField("key", *curfile)
	w.WriteField("KSSAccessKeyId", *accessKey)
	w.WriteField("Policy", string(Base64Encode(json)))
	w.WriteField("Signature", signature)

	f, err := os.Open(*curfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fw, err := w.CreateFormFile("file", "objPost.log")
	if err != nil {
		panic(err)
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	w.Close()

	httpReq, _ := http.NewRequest("POST", urlStr, &buffer)
	httpReq.Header.Add("Content-Type", w.FormDataContentType())
	httpReq.Header.Set("x-kss-newfilename-in-body", "true")
	DoRequest(httpReq)
}

func ObjectPut() {
	urlStr := fmt.Sprintf("http://%s/%s/%s", *addr, *bucketName, *upkey)
	fmt.Printf("\n[%40s:\t%-50s]\n", tracker.Blue("OBJECT-%s-URL", "PUT"), tracker.Yellow(urlStr))

	httpReq, _ := http.NewRequest("PUT", urlStr, strings.NewReader(*content))
	expiresTime := GetDate()
	contentType := "application/octet-stream"
	httpReq.Header.Add("Content-Type", contentType)
	httpReq.Header.Add("date", expiresTime)

	sign := DoSignature("PUT",
		"",
		contentType,
		expiresTime,
		"/"+*bucketName+"/"+*upkey, *secretKey, map[string]string{})

	autoString := fmt.Sprintf("KSS %s:%s", *accessKey, sign)
	httpReq.Header["authorization"] = []string{autoString}

	DoRequest(httpReq)
}

func Object(method, key string) {
	urlStr := fmt.Sprintf("http://%s/%s/%s", *addr, *bucketName, key)
	fmt.Printf("\n[%40s:\t%-50s]\n", tracker.Blue("OBJECT-%s-URL", method), tracker.Yellow(urlStr))

	httpReq, _ := http.NewRequest(method, urlStr, nil)
	expiresTime := GetExpireTime()
	httpReq.Header.Add("date", expiresTime)

	sign := DoSignature(method,
		"",
		"",
		expiresTime,
		"/"+*bucketName+"/"+key, *secretKey, map[string]string{})

	autoString := fmt.Sprintf("KSS %s:%s", *accessKey, sign)
	httpReq.Header["authorization"] = []string{autoString}

	DoRequest(httpReq)
}
