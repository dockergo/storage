package api_test

import (
	"fmt"
	"net/http"

	"github.com/flyaways/tracker"
)

func Bucket(method string) {
	urlStr := fmt.Sprintf("http://%s/%s", *addr, *bucketName)
	expiresTime := GetDate()
	httpReq, _ := http.NewRequest(method, urlStr, nil)
	httpReq.Header.Add("date", expiresTime)
	fmt.Printf("\n[BUCKET-%s-URL:\t%26s]\n", method, tracker.Blue(urlStr))
	sign := DoSignature(method,
		"",
		"",
		expiresTime,
		"/"+*bucketName, *secretKey, map[string]string{})

	autoString := fmt.Sprintf("KSS %s:%s", *accessKey, sign)
	httpReq.Header["authorization"] = []string{autoString}

	DoRequest(httpReq)
}

func Service(method string) {
	urlStr := fmt.Sprintf("http://%s/", *addr)
	expiresTime := GetDate()
	httpReq, _ := http.NewRequest(method, urlStr, nil)
	httpReq.Header.Add("date", expiresTime)
	fmt.Printf("\n[SERVICE-%s-URL:\t%26s]\n", method, tracker.Blue(urlStr))
	sign := DoSignature(method,
		"",
		"",
		expiresTime,
		"/", *secretKey, map[string]string{})

	autoString := fmt.Sprintf("KSS %s:%s", *accessKey, sign)
	httpReq.Header["authorization"] = []string{autoString}

	DoRequest(httpReq)
}
