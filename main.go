package main

import (
	"net/http"
	"os"

	"github.com/bhoriuchi/go-bunyan/bunyan"
	"github.com/gorilla/mux"
)

// globals
var log bunyan.Logger

func main() {
	basePath := getEnv("IOTWEB_BASEPATH", "/")
	staticPath := getEnv("IOTWEB_STATICPATH", "www")

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

	logger := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			staticFields := make(map[string]interface{})
			staticFields["remote"] = r.RemoteAddr
			staticFields["method"] = r.Method
			staticFields["url"] = r.RequestURI

			log.Info(staticFields, "HTTP")
			next.ServeHTTP(w, r)
		})
	}

	r := mux.NewRouter()
	r.Use(logger)

	if basePath != "/" {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, basePath, 301)
		})
		r.PathPrefix(basePath).Handler(http.StripPrefix(basePath, http.FileServer(http.Dir(staticPath))))
	} else {
		r.PathPrefix(basePath).Handler(http.FileServer(http.Dir(staticPath)))
	}

	http.Handle("/", r)
	log.Info("Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getEnv gets an environment variable or sets a default if
// one does not exist.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
