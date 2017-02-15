package storage

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/flyaways/storage/app"
	"github.com/flyaways/storage/util/log"
	"github.com/flyaways/tracker"
	"github.com/gin-gonic/gin"
)

type Policy struct {
	Expiration string
	Conditions map[string]string
}

func (plc *Policy) Marshal() []byte {
	json := `{"expiration": "` + plc.Expiration + `","conditions": [`
	for k, v := range plc.Conditions {
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
	return time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
}

func MakeHmac(key []byte, data []byte) []byte {
	hash := hmac.New(sha1.New, key)
	hash.Write(data)
	return hash.Sum(nil)
}

func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func walkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	//suffix = strings.ToUpper(suffix)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		files = append(files, filename)
		/***
				if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
					files = append(files, filename)
				}
		***/
		return nil
	})
	return files, err
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

	fmt.Printf("[%40s:\t%-50s]\n", tracker.Blue("signature"), tracker.Yellow(string(sign)))

	return sign
}

func objectPost(app *app.App, router *gin.Engine, curfile, bucketName string) {
	urlStr := fmt.Sprintf("/%s", bucketName)
	policy := &Policy{
		Expiration: time.Unix(time.Now().Add(time.Minute*30).Unix(), 0).UTC().Format("2006-01-02T15:04:05.000Z"),
		Conditions: map[string]string{
			"bucket": bucketName,
			"key":    curfile,
		}}

	json := policy.Marshal()
	signature := string(Base64Encode(MakeHmac([]byte(app.Config.Credential.SecretKey), Base64Encode(json))))

	var buffer bytes.Buffer
	w := multipart.NewWriter(&buffer)
	w.WriteField("key", curfile)
	w.WriteField("KSSAccessKeyId", app.Config.Credential.AccessKey)
	w.WriteField("Policy", string(Base64Encode(json)))
	w.WriteField("Signature", signature)

	f, err := os.Open(curfile)
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
	rr := httptest.NewRecorder()
	httpReq.Header.Add("Content-Type", w.FormDataContentType())
	httpReq.Header.Set("x-kss-newfilename-in-body", "true")

	router.ServeHTTP(rr, httpReq)

	if rr.Code == http.StatusOK {
		log.Error("Success: %s", curfile)
	}

}

func initStorage(app *app.App, router *gin.Engine) {
	urlStr := fmt.Sprintf("/%s", "bucketName")
	httpReq, _ := http.NewRequest("PUT", urlStr, nil)
	rr := httptest.NewRecorder()
	expiresTime := GetDate()
	contentType := "application/octet-stream"
	httpReq.Header.Add("Content-Type", contentType)
	httpReq.Header.Add("date", expiresTime)

	sign := DoSignature("PUT",
		"",
		contentType,
		expiresTime,
		"/"+"bucketName"+"/"+"upkey", app.Config.Credential.SecretKey, map[string]string{})

	autoString := fmt.Sprintf("KSS %s:%s", app.Config.Credential.AccessKey, sign)
	httpReq.Header["authorization"] = []string{autoString}
	router.ServeHTTP(rr, httpReq)
	if rr.Code == http.StatusOK {
		log.Error("Success: %s", "curfile")
	}

	files, err := walkDir("../initdir", "*")
	if err != nil {
		log.Error("walkDir error: %s", err.Error())
	}

	for _, curfile := range files {
		objectPost(app, router, curfile, "bucketName")
	}
}
