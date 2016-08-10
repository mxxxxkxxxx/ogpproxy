package ogp

import (
	"golang.org/x/net/html"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/mxxxxkxxxx/ogpproxy/storage/cache"
)

type OgpData struct {
	Title           string `json:"title"`
	Type            string `json:"type"`
	Url             string `json:"url"`
	Image           string `json:"image"`
	Audio           string `json:"audio"`
	Description     string `json:"description"`
	Determiner      string `json:"determiner"`
	Locale          string `json:"locale"`
	LocaleAlternate string `json:"locale_alternate"`
	SiteName        string `json:"site_name"`
	Video           string `json:"video"`
	RequestedUrl    string `json:"-"`
}

func CreateOgpData(root *html.Node, url string) *OgpData {
	ogp := &OgpData{}

	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "meta" {
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
				case "og:type":
					ogp.Type = cont
				case "og:url":
					ogp.Url = cont
				case "og:image":
					ogp.Image = cont
				case "og:audio":
					ogp.Audio = cont
				case "og:description":
					ogp.Description = cont
				case "og:determiner":
					ogp.Determiner = cont
				case "og:locale":
					ogp.Locale = cont
				case "og:locale:alternate":
					ogp.LocaleAlternate = cont
				case "og:site_name":
					ogp.SiteName = cont
				case "og:video":
					ogp.Video = cont
				case "description":
					if len(ogp.Description) == 0 {
						ogp.Description = cont
					}
				}
			} else if n.Data == "title" {
				if len(ogp.Title) == 0 {
					ogp.Title = n.FirstChild.Data
				}
			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(root)
	if (len(ogp.Url) == 0) {
		ogp.Url = url
	}
	ogp.RequestedUrl = url

	return ogp
}

func LoadOgpData(url string) (*OgpData, error) {
	data := &OgpData{}

	cacheHandler := cache.GetHandler()
	buf, err := cacheHandler.Read(url)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to load ogp data: key=[%s]", url)
	}

	// @TODO: check expiration
	err = json.Unmarshal(buf, data)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to convert ogp data from json: key=[%s]", url)
	}

	return data, nil
}

func (o *OgpData) Save() error {
	buf, err := json.Marshal(o)
	if err != nil {
		return errors.Wrapf(err, "Failed to convert ogp data to json: key=[%s]", o.RequestedUrl)
	}

	cacheHandler := cache.GetHandler()
	err = cacheHandler.Write(o.RequestedUrl, buf)
	if err != nil {
		return errors.Wrapf(err, "Failed to save ogp data: key=[%s]", o.RequestedUrl)
	}

	return nil
}