package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/bhoriuchi/go-bunyan/bunyan"
	"github.com/gorilla/mux"
)

type FileInfo struct {
	Name    string      `json:"name"`
	IsDir   bool        `json:"is_dir"`
	Size    int64       `json:"size"`
	ModTime time.Time   `json:"mod_time"`
	Mode    os.FileMode `json:"mode"`
}

func main() {
	basePath := getEnv("IOTWEB_BASEPATH", "/")
	redirectBase := getEnv("IOTWEB_REDIRECT", "yes")
	staticPath := getEnv("IOTWEB_STATICPATH", "www")
	fsApi := getEnv("IOTWEB_FSAPI", "yes")
	fsApiPath := getEnv("IOTWEB_FSAPIPATH", basePath+"fsapi/")
	port := getEnv("IOTWEB_PORT", "8080")

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
	logger, _ := bunyan.CreateLogger(logConfig)

	logHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			staticFields := make(map[string]interface{})
			staticFields["remote"] = r.RemoteAddr
			staticFields["method"] = r.Method
			staticFields["url"] = r.RequestURI

			logger.Info(staticFields, "HTTP")
			next.ServeHTTP(w, r)
		})
	}

	fsApiHandler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		fileFields := make(map[string]interface{})
		path := staticPath + "/" + vars["sub"]

		logger.Info("fsApi: Listing Path: %s", path)

		files, err := ioutil.ReadDir(path)
		if err != nil {
			logger.Info("fsApi: Path Error: %s", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{ \"error\": \"Not Found\"}"))
			return
		}

		for _, f := range files {
			fileFields[f.Name()] = FileInfo{
				Name:    f.Name(),
				IsDir:   f.IsDir(),
				Size:    f.Size(),
				ModTime: f.ModTime(),
				Mode:    f.Mode(),
			}
		}

		js, err := json.Marshal(fileFields)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	allowHeaders := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With, Accept, Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")

			next.ServeHTTP(w, r)
		})
	}

	r := mux.NewRouter()
	r.Use(logHandler)
	r.Use(allowHeaders)

	if fsApi == "yes" {
		r.HandleFunc(fsApiPath+"{sub:.*}", fsApiHandler)
	}

	if basePath != "/" && redirectBase == "yes" {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, basePath, 301)
		})
	}

	r.PathPrefix(basePath).Handler(http.StripPrefix(basePath, http.FileServer(http.Dir(staticPath))))

	http.Handle("/", r)
	logger.Info("Listening on Port " + port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
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
