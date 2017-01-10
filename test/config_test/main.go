package main

import (
	"flag"
	"os"

	"github.com/flyaways/storage/agent/config"
	"github.com/flyaways/storage/agent/storage/kdfs"
	"github.com/flyaways/storage/agent/storage/nfs"
	"github.com/flyaways/storage/agent/storage/posix"
	"github.com/flyaways/storage/agent/storage/s3"
	"github.com/flyaways/storage/agent/storage/swift"
	"github.com/flyaways/storage/agent/util/log"
)

var configFile = flag.String("config", "agent.toml", "agent config file")

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

	storagenfs := nfs.New(cfg)
	if storagenfs != nil {
		log.Error("storagenfs error")
		panic(storagenfs)
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

	storagekdfs := kdfs.New(cfg)
	if storagekdfs != nil {
		log.Error("storagekdfs error")
		panic(storagekdfs)
	}

	storageswift := swift.New(cfg)
	if storageswift != nil {
		log.Error("storageswift error")
		panic(storageswift)
	}

}
