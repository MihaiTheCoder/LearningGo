package main

import (
	"net/http"
	"strings"
	"os"
	"bufio"
)

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path

	f, err := os.Open(path)

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
		return
	}
	contentType := getContentType(path)
	w.Header().Add("Content Type", contentType)
	bufferedReader := bufio.NewReader(f)
	bufferedReader.WriteTo(w)
}

func getContentType(path string) (contentType string) {
	contentType = defaultContentType

	lastIndexOfDot := strings.LastIndex(path, ".")

	if lastIndexOfDot < 0 {
		return
	}

	extension := path[lastIndexOfDot:]

	if tempContentType, ok := extensionToContentType[extension]; ok {
		contentType = tempContentType
	}
	return
}

var extensionToContentType = map[string]string{
	".css":  "text/css",
	".html": "text/html",
	".js":   "application/javascript",
	".png":  "image/png",
	".mp4":   "video/mp4",
}

const defaultContentType = "text/plain"