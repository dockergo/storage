package api_test

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/flyaways/tracker"
)

func DoRequest(httpReq *http.Request) {
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		MaxIdleConnsPerHost: 180,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	httpRep, err := client.Do(httpReq)
	if err != nil {
		fmt.Printf("\n%s\n", err.Error())
	}

	fmt.Printf("[%40s:\t%-50s]\n", tracker.Magenta("status"), tracker.Red("%d", httpRep.StatusCode))

	body, err := ioutil.ReadAll(httpRep.Body)
	if err != nil {
		fmt.Printf("\n[%s]\n", tracker.Red(err.Error()))
	}

	if httpRep.Header.Get(Newfilename) != "" {
		*newName = httpRep.Header.Get(Newfilename)
	}

	for key, value := range httpRep.Header {
		for _, values := range value {
			fmt.Printf("[%40s:\t%-50v]\n", tracker.Magenta(key), tracker.Green(values))
		}
	}

	fmt.Printf("[%40s:\t%-50s]\n", tracker.Magenta("bodySize"), tracker.Green("%d", len(body)))
	fmt.Printf("%40s:\n%s\n", tracker.Magenta("body"), tracker.Cyan(string(body)))

}
