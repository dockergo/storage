//Website: https://github.com/flyaways
//Modifier: Flyaway
//Date 27/09/2016 13:22 Beijing
//
package swift

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/config"
	errs "github.com/flyaways/storage/agent/errors"
	"github.com/flyaways/storage/agent/result"
	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util/log"
)

type Swift struct {
	adapter.StorageAdapter
	config      *config.Config
	httpClient  *http.Client
	authToken   string
	authAccount string
}

func New(config *config.Config) *Swift {
	swt := new(Swift)
	swt.config = config
	swt.Name = "swift"
	swt.httpClient = &http.Client{
		Timeout: time.Duration(time.Minute),
	}

	err := swt.Auth()
	if err != nil {
		log.Error("[%s:%s]", swt.Name, err.Error())
		panic("swift auth failed")
	}

	return swt
}

func doRequest(method, url, token string, body io.Reader, client *http.Client) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Auth-Token", token)
	resp, err = client.Do(req)
	return
}

func (swt *Swift) Auth() error {
	url, auxjson, err := buildAuthReq(swt.config)
	if err != nil {
		log.Error("[%s:%s]", swt.Name, err.Error())
		return err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(auxjson))
	if err != nil {
		log.Error("[%s:%s]", swt.Name, err.Error())
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpResp, err := swt.httpClient.Do(httpReq)
	if err != nil {
		log.Error("[%s:%s]", swt.Name, err.Error())
		return err
	}

	swt.authToken = httpResp.Header.Get("X-Subject-Token")

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Error("[%s:%s]", swt.Name, err.Error())
		return err
	}

	defer httpResp.Body.Close()
	var r TokenR
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Error("[%s:%s]", swt.Name, err.Error())
		return err
	}
	swt.authAccount = "AUTH_" + r.Token.ProjectT.Id
	return nil

}

func buildUrl(addr, account, bucket, key string) string {
	return fmt.Sprintf("%s/%s", buildBucketUrl(addr, account, bucket), key)
}

func buildBucketUrl(addr, account, bucket string) string {
	return fmt.Sprintf("%s/%s/%s", addr, account, bucket)
}

func buildAuthReq(config *config.Config) (string, []byte, error) {
	var r AuthR
	r.Auth.Scope.Project.Domain.Name = config.Storage.Swift.TenantName
	r.Auth.Scope.Project.Name = config.Storage.Swift.ProjName
	r.Auth.Identity.Password.User.Domain.Name = config.Storage.Swift.TenantName
	r.Auth.Identity.Password.User.Name = config.Storage.Swift.UserName
	r.Auth.Identity.Password.User.Password = config.Storage.Swift.PassWord
	r.Auth.Identity.Methods = []string{0: "password"}

	bytes, err := json.Marshal(&r)
	if err != nil {
		log.Error("[%s:%s]", config.Storage.Type, err.Error())
	}
	return fmt.Sprintf("%s", config.Storage.Swift.AuthURL), bytes, err
}

func (swt *Swift) request(data io.Reader, method, url string, res *result.Result, ctx *gin.Context) {
	httpResponse, err := doRequest(method, url, swt.authToken, data, swt.httpClient)
	if err != nil {
		err = swt.Auth()
		if err == nil {
			httpResponse, err = doRequest(method, url, swt.authToken, data, swt.httpClient)
		}
	}

	if err != nil {
		log.Error("[%s:%s]", swt.Name, err.Error())
		res.Error(errs.InvalidArgument)
	}

	ctx.Status(httpResponse.StatusCode)
}
