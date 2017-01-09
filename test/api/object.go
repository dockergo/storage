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

	fmt.Printf("\nOBJECT--POST----: %s\n", tracker.Blue(urlStr))
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

	f, err := os.Open(*curfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w.WriteField("key", *curfile)
	w.WriteField("KSSAccessKeyId", *accessKey)
	w.WriteField("Policy", string(Base64Encode(json)))
	w.WriteField("Signature", signature)
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
	urlStr := fmt.Sprintf("http://%s/%s/%s", *addr, *bucketName, *key)
	expiresTime := GetDate()
	contentType := "application/octet-stream"

	fmt.Printf("\nOBJECT--PUT----: %s\n", tracker.Blue(urlStr))
	httpReq, _ := http.NewRequest("PUT", urlStr, strings.NewReader(*content))
	httpReq.Header.Add("Content-Type", contentType)
	httpReq.Header.Add("date", expiresTime)

	sign := DoSignature("PUT",
		"",
		contentType,
		expiresTime,
		"/"+*bucketName+"/"+*key, *secretKey, map[string]string{})

	autoString := fmt.Sprintf("KSS %s:%s", *accessKey, sign)
	httpReq.Header["authorization"] = []string{autoString}

	DoRequest(httpReq)
}

func Object(method, key string) {
	urlStr := fmt.Sprintf("http://%s/%s/%s", *addr, *bucketName, key)
	expiresTime := GetExpireTime()

	fmt.Printf("\nOBJECT--%s----: %s\n", method, tracker.Blue(urlStr))

	var body io.Reader
	body = nil
	if method == "PUT" {
		body = strings.NewReader(*content)
	}

	httpReq, _ := http.NewRequest(method, urlStr, body)
	httpReq.Header.Add("date", expiresTime)

	if method == "PUT" {
		contentType := "application/octet-stream"
		httpReq.Header.Add("Content-Type", contentType)

	}

	sign := DoSignature(method,
		"",
		"",
		expiresTime,
		"/"+*bucketName+"/"+key, *secretKey, map[string]string{})

	autoString := fmt.Sprintf("KSS %s:%s", *accessKey, sign)
	httpReq.Header["authorization"] = []string{autoString}

	DoRequest(httpReq)
}
