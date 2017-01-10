package main

import (
	"flag"

	test "github.com/flyaways/storage/test/api"
)

func RuningCLI() {
	flag.Parse()
	test.TouchFile()

	if *test.Types == "bucket" {
		test.Bucket(*test.Method)
	} else if *test.Types == "object" {
		test.MatchObject(*test.Method)
	} else {
		test.Service("GET")
	}
}

func main() {
	RuningCLI()
}
