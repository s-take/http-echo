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
	mux.HandleFunc("/", index)
	mux.HandleFunc("/dump", dumprequest)
	mux.HandleFunc("/wait", wait)
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is echo service\n")
}

func dumprequest(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	io.WriteString(w, "This is echo service\n")
	io.WriteString(w, "===DumpRequest===\n")
	io.WriteString(w, string(dump))
}

func wait(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	io.WriteString(w, "Waited 10 seconds \n")
}
