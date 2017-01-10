package api_test

import (
	"fmt"
	"net/http"

	"github.com/flyaways/tracker"
)

func Bucket(method string) {
	urlStr := fmt.Sprintf("http://%s/%s", *addr, *bucketName)
	fmt.Printf("\n[%40s:\t%-50s]\n", tracker.Blue("BUCKET-%s-URL", method), tracker.Yellow(urlStr))

	httpReq, _ := http.NewRequest(method, urlStr, nil)
	expiresTime := GetDate()
	httpReq.Header.Add("date", expiresTime)

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
	fmt.Printf("\n[%40s:\t%-50s]\n", tracker.Blue("SERVICE-%s-URL", method), tracker.Yellow(urlStr))

	httpReq, _ := http.NewRequest(method, urlStr, nil)
	expiresTime := GetDate()
	httpReq.Header.Add("date", expiresTime)
	sign := DoSignature(method,
		"",
		"",
		expiresTime,
		"/", *secretKey, map[string]string{})

	autoString := fmt.Sprintf("KSS %s:%s", *accessKey, sign)
	httpReq.Header["authorization"] = []string{autoString}

	DoRequest(httpReq)
}
