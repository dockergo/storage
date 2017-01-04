package api_test

func MatchObject(method string) {
	switch method {
	case "GET":
		Object(method, *newName)
	case "DELETE":
		Object(method, *newName)
	case "PUT":
		ObjectPut()
	case "POST":
		ObjectPost()
	case "HEAD":
		Object(method, *newName)
	default:
		return
	}
}
