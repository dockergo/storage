package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/flyaways/storage/agent"
	"github.com/flyaways/storage/agent/config"
)

var configFile = flag.String("config", "agent.toml", "agent config file")

func main() {

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(*configFile) == 0 {
		println("no config set")
		os.Exit(1)
	}
	cfg, err := config.ParseConfig(*configFile)

	if err != nil {
		println("Parse config:", err.Error())
		os.Exit(1)
	}

	server, err := agent.NewServer(cfg)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	server.Run()
	//server.RunTLS(addr string, certFile string, keyFile string)

}
