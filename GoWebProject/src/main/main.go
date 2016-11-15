package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", serveHTTP)
	http.ListenAndServe(":8000", nil)
}

func serveHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	templates := template.New("xx")
	templates.New("test").Parse(doc);
	templates.New("header").Parse(header)
	templates.New("footer").Parse(footer)
	context := newContext("the title", "Lemon", "Orange", "Apple")
	templates.Lookup("test").Execute(w, context)
}

const doc = `
{{template "header" .Title}}
	<body>
		<h1>List of Fruits</h1>
		<ul>
		{{range .Fruits}}
			<li>{{.}}</li>
		{{end}}
		</ul>
	</body>
{{template "footer"}}
`
const header = `
<!DOCTYPE html>
<html>
	<head><title>{{.}}</title></head>
`
const footer = `
</html>
`

func newContext(title string, fruits ...string) (context Context)  {
	context = Context{
		Fruits: fruits,
		Title: title}
	return
}
type Context struct {
	Fruits []string
	Title string
}