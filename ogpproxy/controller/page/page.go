package page

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
	"ogpproxy/ogpproxy/controller"
	"ogpproxy/ogpproxy/storage/cache"
	"ogpproxy/ogpproxy/ogp"
	"ogpproxy/ogpproxy/console"
)

func Get(w http.ResponseWriter, r *http.Request) {
	res := &controller.Response{Writer: w}

	if r.Method != "GET" {
		res.WriteError("GET is only supported.")
		return
	}

	values := r.URL.Query()
	url := ""
	if v, ok := values["url"]; ok {
		url = v[len(v) - 1]
	}
	if len(url) == 0 {
		res.WriteError("You must specify a variable named url.")
		return
	}

	cacheHandler := cache.GetHandler()

	if cache, err := cacheHandler.Read(url); err == nil {
		console.Info("GET " + url + " from cache")
		res.Ogp = cache
	} else {
		console.Debug("Failed to read from cache: err=[" + err.Error() + "]")
		console.Info("GET " + url + " from remote")

		destRes, err := http.Get(url)
		if err != nil {
			res.WriteError(fmt.Sprintf("Failed to do GET %s", url))
			return
		}
		defer destRes.Body.Close()

		doc, err := html.Parse(destRes.Body)
		if err != nil {
			res.WriteError(fmt.Sprintf("Failed to parse %s", url))
			return
		}

		res.Ogp = ogp.CreateOgpData(doc, url)
		go func() {
			console.Debug("Trying to write cache... : url=[" + url + "]")
			err = cacheHandler.Write(res.Ogp)
			if err != nil {
				console.Error("Failed to write cache: url=[" + url + "], err=[" + err.Error() + "]")
			}
		}()
	}

	res.Write()
}

