package main

import (
	"flag"
	"github.com/Flyaways/tracker"
	test "github.com/flyaways/storage/test/api"
)

func RuningCLI() {
	flag.Parse()
	test.TouchFile()

	if *test.Types == "bucket" {
		test.Bucket(*test.Method)
	} else {
		test.MatchObject(*test.Method)
	}
}

func main() {
	tracker.Tracker("notice", test.Common, "MATCH TESTING")
	RuningCLI()
	tracker.Tracker("notice", test.Common, "MATCH TESTING OVER!")
}
