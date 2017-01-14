package api_test

func MatchObject(method string) {
	switch method {
	case "GET":
		Object(method, *reqkey)
	case "DELETE":
		Object(method, *reqkey)
	case "PUT":
		ObjectPut()
	case "POST":
		ObjectPost()
	case "HEAD":
		Object(method, *reqkey)
	case "OPTIONS":
		Object(method, *reqkey)

	default:

		return
	}
}
