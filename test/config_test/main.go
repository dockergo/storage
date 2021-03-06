package main

import (
	"flag"
	"os"

	"github.com/flyaways/storage/config"

	"github.com/flyaways/storage/types/oss"
	"github.com/flyaways/storage/types/posix"
	"github.com/flyaways/storage/types/qiniu"
	"github.com/flyaways/storage/types/s3"
	"github.com/flyaways/storage/types/swift"

	"github.com/flyaways/storage/util/log"
)

var configFile = flag.String("config", "storage.toml", "storage config file")

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("storage error: %s", r)
			os.Exit(1)
		}
	}()

	if len(*configFile) == 0 {
		log.Error("no config set")
		os.Exit(1)
	}
	cfg, err := config.ParseConfig(*configFile)

	if err != nil {
		log.Error("Parse config:", err.Error())
		os.Exit(1)
	}

	storageoss := oss.New(cfg)
	if storageoss != nil {
		log.Error("storagenfs error")
		panic(storageoss)
	}
	storageposix := posix.New(cfg)
	if storageposix != nil {
		log.Error("storageposix error")
		panic(storageposix)
	}

	storages3 := s3.New(cfg)
	if storages3 != nil {
		log.Error("storages3 error")
		panic(storages3)
	}

	storageswift := swift.New(cfg)
	if storageswift != nil {
		log.Error("storageswift error")
		panic(storageswift)
	}

	storageqiniu := qiniu.New(cfg)
	if storageqiniu != nil {
		log.Error("storageqiniu error")
		panic(storageqiniu)
	}

}
