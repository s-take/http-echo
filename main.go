package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", dump)
	mux.HandleFunc("/slow", slow)
	mux.HandleFunc("/error", err)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func dump(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	io.WriteString(w, "This is echo service\n")
	io.WriteString(w, "===DumpRequest===\n")
	io.WriteString(w, string(dump))
}

func slow(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is echo service\n")
	time.Sleep(10 * time.Second)
	io.WriteString(w, "Waited 10 seconds \n")
}

func err(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)

	io.WriteString(w, "This is echo service\n")
	io.WriteString(w, "Error!! \n")
}
