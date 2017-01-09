package main

import (
	"flag"

	test "github.com/flyaways/storage/test/api"
	"github.com/flyaways/tracker"
)

var buketmethodlist = [...]string{0: "GET", 1: "HEAD"}

var objectmethodlist = [...]string{0: "PUT", 1: "GET", 2: "DELETE", 3: "POST", 4: "HEAD"}

func RuningOneKeyTest() {
	flag.Parse()
	test.TouchFile()

	tracker.Tracker("notice", test.Common, "BUCKET CREATE")
	test.Bucket("PUT")

	tracker.Tracker("notice", test.Common, "OBJECT TESTING")
	for _, arg := range objectmethodlist {
		test.MatchObject(arg)
	}

	tracker.Tracker("notice", test.Common, "SERVICE TESTING")
	test.Service("GET")

	tracker.Tracker("notice", test.Common, "BUCKET TESTING")
	for _, arg := range buketmethodlist {
		test.Bucket(arg)
	}

	tracker.Tracker("notice", test.Common, "OBJECT DELETE")
	test.MatchObject("DELETE")

	tracker.Tracker("notice", test.Common, "BUCKET DELETE")
	test.Bucket("DELETE")
}

func main() {
	tracker.Tracker("notice", test.Common, "GROUP TESTING")
	RuningOneKeyTest()
	tracker.Tracker("notice", test.Common, "STORAGE TESTING OVER!")
}
