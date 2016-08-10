package page

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
	"github.com/mxxxxkxxxx/ogpproxy/controller"
	"github.com/mxxxxkxxxx/ogpproxy/ogp"
	"github.com/mxxxxkxxxx/ogpproxy/console"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/japanese"
	"io"
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

	if cache, err := ogp.LoadOgpData(url); err == nil {
		console.Info("GET " + url + " from cache")
		console.Debug(fmt.Sprintf("ogp(cached): %+v", cache))
		res.Ogp = cache
	} else {
		console.Debug("Failed to read from cache: err=[" + err.Error() + "]")
		console.Info("GET " + url + " from remote")


		doc, err := getHtmlFromRemote(url)
		if err != nil {
			res.WriteError(fmt.Sprintf("Failed to do GET %s", url))
			console.Error("Failed to do GET: url=[" + url + "], err=[" + err.Error() + "]")
			return
		}

		res.Ogp = ogp.CreateOgpData(doc, url)
		console.Debug(fmt.Sprintf("ogp: %+v", res.Ogp))

		go func() {
			console.Debug("Trying to write cache... : url=[" + url + "]")
			err = res.Ogp.Save()
			if err != nil {
				console.Error("Failed to write cache: url=[" + url + "], err=[" + err.Error() + "]")
			}
		}()
	}

	res.Write()
}

func getHtmlFromRemote(url string) (*html.Node, error) {
	destRes, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer destRes.Body.Close()

	console.Debug(fmt.Sprintf("response headers: %+v", destRes.Header))

	charSet := ""
	if v, ok := destRes.Header["Content-Type"]; ok {
		cTypeVals := strings.Split(v[0], ";")
		if len(cTypeVals) == 2 {
			charSetVals := strings.Split(strings.TrimSpace(cTypeVals[1]), "=")
			if len(charSetVals) == 2 {
				charSet = charSetVals[1]
			}
		}
	}
	console.Debug("charset: " + charSet)

	var reader io.Reader
	decoderFound := false
	if len(charSet) > 0 {
		switch (strings.ToUpper(charSet)) {
		case "EUC_JP":
			fallthrough
		case "EUC-JP":
			reader = transform.NewReader(destRes.Body, japanese.EUCJP.NewDecoder())
			decoderFound = true
		case "SHIFT_JIS":
			fallthrough
		case "SHIFT-JIS":
			reader = transform.NewReader(destRes.Body, japanese.ShiftJIS.NewDecoder())
			decoderFound = true
		}
	}

	var doc *html.Node;
	if decoderFound {
		doc, err = html.Parse(reader)
	} else {
		doc, err = html.Parse(destRes.Body)
	}

	return doc, err
}