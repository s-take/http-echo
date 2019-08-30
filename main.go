package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", dumprequest)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func dumprequest(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	io.WriteString(w, "This is echo service\n")
	io.WriteString(w, "===DumpRequest===\n")
	io.WriteString(w, string(dump))
}
