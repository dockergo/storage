package main

import (
	"flag"
	"github.com/Flyaways/tracker"
	test "github.com/flyaways/storage/test/api"
)

var buketmethodlist = [...]string{0: "PUT", 1: "HEAD"}

var objectmethodlist = [...]string{0: "PUT", 1: "GET", 2: "DELETE", 3: "POST", 4: "HEAD", 5: "DELETE"}

func RuningOneKeyTest() {
	flag.Parse()
	test.TouchFile()

	tracker.Tracker("notice", test.Common, "BUCKET TESTING BEGIN")
	for _, arg := range buketmethodlist {
		test.Bucket(arg)
	}
	tracker.Tracker("notice", test.Common, "OBJECT TESTING BEGIN")

	for _, arg := range objectmethodlist {
		test.MatchObject(arg)
	}
	tracker.Tracker("notice", test.Common, "BUCKET DELETE BEGIN")
	test.Bucket("DELETE")
}

func main() {
	tracker.Tracker("notice", test.Common, "GROUP TESTING")
	RuningOneKeyTest()
	tracker.Tracker("notice", test.Common, "STORAGE TESTING OVER!")
}
