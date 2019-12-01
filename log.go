package main

import (
	"bytes"
	"log"
	"net/http"
	"text/template"
)

type LineOfLog struct {
	RemoteAddr  string
	ContentType string
	Path        string
	Query       string
	Method      string
	Body        string
}

var TemplateOfLog = `
Remote address:   {{.RemoteAddr}}
Content-Type:     {{.ContentType}}
HTTP method:      {{.Method}}

path:
{{.Path}}

query string:
{{.Query}}

body:             
{{.Body}}

`

func AccessLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(r.Body)
		body := bufbody.String()

		line := LineOfLog{
			r.RemoteAddr,
			r.Header.Get("Content-Type"),
			r.URL.Path,
			r.URL.RawQuery,
			r.Method, body,
		}
		tmpl, err := template.New("line").Parse(TemplateOfLog)
		if err != nil {
			panic(err)
		}

		bufline := new(bytes.Buffer)
		err = tmpl.Execute(bufline, line)
		if err != nil {
			panic(err)
		}

		log.Printf(bufline.String())
		handler.ServeHTTP(w, r)
	})
}
