package api_test

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/Flyaways/tracker"
)

func DoRequest(httpReq *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\n%s\n", r)
		}
	}()
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
		panic(err)
	}
	tmp := fmt.Sprintf("%d", httpRep.StatusCode)
	fmt.Printf("response-----status:  %s\n", tracker.Red(tmp))
	body, err := ioutil.ReadAll(httpRep.Body)

	if httpRep.Header.Get(Newfilename) != "" {
		*newName = httpRep.Header.Get(Newfilename)
	}

	fmt.Printf("response-----Newfilename:  %s\n", tracker.Red(*newName))
	lens := fmt.Sprintf("%d", len(body))
	fmt.Printf("response-----body:   %s, %s\n", tracker.Yellow(string(body)), tracker.Magenta(lens))
	if err != nil {
		fmt.Printf("\n%s\n", err.Error())
		panic(err)
	} else {
	}
}

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
