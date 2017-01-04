package api_test

import (
	"fmt"
	"net/http"

	"github.com/Flyaways/tracker"
)

func Bucket(method string) {
	urlStr := fmt.Sprintf("http://%s/%s", *addr, *bucketName)
	expiresTime := GetDate()
	httpReq, _ := http.NewRequest(method, urlStr, nil)
	httpReq.Header.Add("date", expiresTime)
	fmt.Printf("\nBUCKET--%s----: %s\n", method, tracker.Blue(urlStr))
	sign := DoSignature(method,
		"",
		"",
		expiresTime,
		"/"+*bucketName, *secretKey, map[string]string{})

	autoString := fmt.Sprintf("KSS %s:%s", *accessKey, sign)
	httpReq.Header["authorization"] = []string{autoString}

	DoRequest(httpReq)
}
