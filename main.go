package main

import (
	"net/http"
	"os"

	"github.com/bhoriuchi/go-bunyan/bunyan"
)

// globals
var log bunyan.Logger

func main() {
	staticFields := make(map[string]interface{})
	staticFields["name"] = "iotweb"

	// configure the logger
	logConfig := bunyan.Config{
		Name:         "iotweb",
		Stream:       os.Stdout,
		Level:        bunyan.LogLevelDebug,
		StaticFields: staticFields,
	}

	// Create a logger
	// see https://github.com/bhoriuchi/go-bunyan
	log, _ = bunyan.CreateLogger(logConfig)

	http.HandleFunc("/", static)

	log.Info("Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func static(w http.ResponseWriter, r *http.Request) {
	staticPath := "./www"

	staticFields := make(map[string]interface{})
	staticFields["remote"] = r.RemoteAddr
	staticFields["method"] = r.Method
	staticFields["url"] = r.RequestURI

	log.Info(staticFields, "STATIC")

	fs := http.FileServer(http.Dir(staticPath))
	fs.ServeHTTP(w, r)
}
