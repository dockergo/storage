package kdfs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/config"
	"github.com/flyaways/storage/agent/result"
	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util/log"
)

type Kdfs struct {
	adapter.StorageAdapter
	config     *config.Config
	httpClient *http.Client
}

func New(config *config.Config) *Kdfs {
	Kdfs := new(Kdfs)
	Kdfs.config = config
	Kdfs.Name = "kdfs"
	Kdfs.httpClient = &http.Client{
		Timeout: time.Duration(time.Minute),
	}

	return Kdfs
}

func doRequest(method string, url string, body io.Reader, client *http.Client) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func buildUrl(addr, account, bucket, key string) string {
	return fmt.Sprintf("%s/%s", buildBucketUrl(addr, account, bucket), key)
}

func buildBucketUrl(addr, account, bucket string) string {
	return fmt.Sprintf("%s/%s/%s", addr, account, bucket)
}

func (kfs *Kdfs) bucketrequest(data io.Reader, method, url string, res *result.Result, ctx *gin.Context) {
	httpResp, err := doRequest(method, url, data, kfs.httpClient)
	if err != nil {
		log.Error("[%s:%s]", kfs.Name, err.Error())
		res.Error(err)
	}

	ctx.Status(httpResp.StatusCode)
	content, _ := ioutil.ReadAll(httpResp.Body)
	ctx.XML(httpResp.StatusCode, content)
}

func (kfs *Kdfs) request(data io.Reader, method, url string, res *result.Result, ctx *gin.Context) {
	httpResp, err := doRequest(method, url, data, kfs.httpClient)
	if err != nil {
		log.Error("[%s:%s]", kfs.Name, err.Error())
		res.Error(err)
	}

	ctx.Status(httpResp.StatusCode)
}
