package page

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
	"ogpproxy/ogpproxy"
	"ogpproxy/ogpproxy/storage/cache"
)

func extractOgpData(root *html.Node) *ogpproxy.OgpData {
	ogp := &ogpproxy.OgpData{}

	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			var prop, cont string
			for _, attr := range n.Attr {
				switch attr.Key {
				case "property":
					prop = attr.Val
				case "content":
					cont = attr.Val
				}
			}

			if len(prop) == 0 || len(cont) == 0 {
				return
			}

			switch prop {
			case "og:title":
				ogp.Title = cont
			case "og:description":
				ogp.Description = cont
			case "og:url":
				ogp.Url = cont
			case "og:image":
				ogp.Image = cont
			case "og:site_name":
				ogp.SiteName = cont
			case "og:locale":
				ogp.Locale = cont
			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(root)

	return ogp
}

func Get(w http.ResponseWriter, r *http.Request) {
	res := &ogpproxy.Response{Writer: w}

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
		fmt.Printf("GET %s from cache\n", url)
		res.Ogp = cache
	} else {
		fmt.Printf("Failed to read from cache: err=[%s]\n", err)
		fmt.Printf("GET %s from remote\n", url)

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

		res.Ogp = extractOgpData(doc)
		go cacheHandler.Write(res.Ogp)
	}

	res.Write()
}

