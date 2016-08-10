package main

import (
	"fmt"
	"os"
	"net/http"
	"ogpproxy/ogpproxy/controller/page"
	"ogpproxy/ogpproxy/config"
	"ogpproxy/ogpproxy/console"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	config := config.GetConfig()
	console.Debug(fmt.Sprintf("config: %+v", config))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusNotFound)
	})
	http.HandleFunc("/", page.Get)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), nil)
	if (err != nil) {
		console.Error("An error occured. err=[" + err.Error() + "]")
		return 1
	}

	return 0
}