package main

import (
	"fmt"
	"flag"
	"os"
	"net/http"
	"ogpproxy/ogpproxy/controllers/page"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	var port int
	flag.IntVar(&port, "port", 8080, "Port number to listen")

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusNotFound)
	})
	http.HandleFunc("/", page.Get)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if (err != nil) {
		fmt.Printf("An error occured. err=[%s]", err)
		return 1
	}

	return 0
}