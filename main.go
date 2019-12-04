package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var requests int64 = 0

// increments the number of requests and returns the new value
func incRequests() int64 {
	return atomic.AddInt64(&requests, 1)
}

// returns the current value
func getRequests() int64 {
	return atomic.LoadInt64(&requests)
}

// increments the number of requests and returns the new value
func clearRequests() error {
	requests = 0
	return nil
}

var logger *zap.Logger

func init() {

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts1",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	// config.OutputPaths = []string{"./log/error.log"}
	config.OutputPaths = []string{"stdout"}
	logger, _ = config.Build()

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", dump)
	mux.HandleFunc("/clearrequests", clearrequests)
	mux.HandleFunc("/slow", slow)
	mux.HandleFunc("/error", err)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func dump(w http.ResponseWriter, r *http.Request) {
	incRequests()
	cnt := getRequests()
	logger.Info("test", zap.String("url", r.URL.Path), zap.Int64("count", cnt))
	dump, _ := httputil.DumpRequest(r, true)
	io.WriteString(w, "This is echo service\n")
	io.WriteString(w, "===DumpRequest===\n")
	io.WriteString(w, string(dump))
}

func clearrequests(w http.ResponseWriter, r *http.Request) {
	clearRequests()
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
