package main

import (
	"net/http"
	"text/template"
	"fmt"
)

func main() {
	http.HandleFunc("/", serveHTTP)
	http.ListenAndServe(":8000", nil)
}

func serveHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("test").Parse(doc)

	if err != nil {
		fmt.Println(err)
		return
	}
	context := newContext()
	tmpl.Execute(w, context)
}

const doc = `
<!DOCTYPE html>
<html>
	<head><title>Example Title</title></head>
	<body>
		<h1> Hello {{.Message}}</h1>
	</body>
</html>
`

func newContext() (context Context)  {
	context = Context{"the message"}
	return
}
type Context struct {
	Message string
}