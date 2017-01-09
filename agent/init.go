package agent

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/flyaways/storage/agent/app"
	"github.com/flyaways/storage/agent/config"
	"github.com/flyaways/storage/agent/constant"
	"github.com/flyaways/storage/agent/storage/adapter"
	"github.com/flyaways/storage/agent/util"
	"github.com/flyaways/storage/agent/util/log"
	"github.com/flyaways/tracker"
)

func StorageInit(app *app.App) {
	initBucket(app.Storager, app.Config)
	initObject(app.Storager, app.Config)
}

func walkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		files = append(files, filename)
		log.Info("%s", tracker.Blue(filename))
		return nil
	})
	return files, err
}

func Policy(bucket, key string, raw []byte) (finalkey string) {
	finalkey = util.GetSha1Hex(raw)
	if strings.HasPrefix(bucket, "bk") && len(bucket) >= 3 {
		policy := bucket[2:3]
		base, err := strconv.ParseInt(policy, 16, 64)
		if err == nil {
			if (base%constant.Key - base%constant.Auth) == 4 {
				finalkey = key
			}
		}
	}
	return
}

func initObject(Storage adapter.Storager, cfg *config.Config) {
	files, err := walkDir(cfg.Init.UploadDir, "")
	if err != nil {
		log.Error("[walkDir\t%s]", files)
		panic(err)
	}

	var data adapter.InitData
	for _, path := range files {
		for _, bucket := range cfg.Init.Buckets {
			rawdata, err := ioutil.ReadFile(path)
			if err != nil {
				log.Error("[ReadFile\t%s]", err.Error())
			}
			finalkey := Policy(bucket, path, rawdata)
			data.RawData = rawdata
			data.Key = finalkey
			data.Bucket = bucket
			Storage.InitObject(&data)

			log.Info("[%19s]%s%-32s", tracker.Blue(bucket), tracker.Green(" <=======> "), filepath.Base(finalkey))

		}
	}
}

func initBucket(Storage adapter.Storager, cfg *config.Config) {
	var data adapter.InitData
	for _, bucket := range cfg.Init.Buckets {
		data.Bucket = bucket
		if err := Storage.InitBucket(&data); err != nil {
			log.Error("[error:%s]", err.Error())
			os.Exit(1)
		}
		log.Info("[%19s]", tracker.Blue(bucket))
	}
}
