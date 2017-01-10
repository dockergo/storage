package main

import (
	"flag"

	test "github.com/flyaways/storage/test/api"
)

var buketmethodlist = [...]string{0: "GET", 1: "HEAD"}

var objectmethodlist = [...]string{0: "PUT", 1: "GET", 2: "DELETE", 3: "POST", 4: "HEAD"}

func RuningOneKeyTest() {
	flag.Parse()
	test.TouchFile()

	test.Bucket("PUT")

	for _, arg := range objectmethodlist {
		test.MatchObject(arg)
	}
	test.Service("GET")

	for _, arg := range buketmethodlist {
		test.Bucket(arg)
	}

	test.MatchObject("DELETE")

	test.Bucket("DELETE")
}

func main() {
	RuningOneKeyTest()
}
