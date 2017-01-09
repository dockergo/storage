package main

import (
	"flag"

	test "github.com/flyaways/storage/test/api"
	"github.com/flyaways/tracker"
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
	tracker.Tracker("notice", test.Common, "MATCH TESTING")
	RuningCLI()
	tracker.Tracker("notice", test.Common, "MATCH TESTING OVER!")
}
