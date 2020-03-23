package rest

import "net/http"

func initPosts() []*post {
	return []*post{
		{ID: "1", Title: "My first post", Body: "This is the content of my first post"},
		{ID: "2", Title: "My second post", Body: "This is the content of my second post"},
	}
}

// INFO: no need to test this function, it should be already tested by "net/http" library
func setJsonContentType(writer http.ResponseWriter) {
	writer.Header().Set(contentTypeHeaderKey, contentTypeHeaderValue)
}

// INFO: no need to test this function, it should be already tested by "net/http" library
func setStatusAccepted(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusAccepted)
}
