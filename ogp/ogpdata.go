package ogp

import (
	"golang.org/x/net/html"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/mxxxxkxxxx/ogpproxy/storage/cache"
	"strconv"
	"github.com/mxxxxkxxxx/ogpproxy/console"
)

type OgpData struct {
	Title           []string        `json:"title"`
	Type            []string        `json:"type"`
	Url             []string        `json:"url"`
	Image           []OgpImageData  `json:"image"`
	Video           []OgpVideoData  `json:"video"`
	Audio           []OgpAudioData  `json:"audio"`
	Description     []string        `json:"description"`
	Determiner      []string        `json:"determiner"`
	Locale          []OgpLocaleData `json:"locale"`
	SiteName        []string        `json:"site_name"`
	RequestedUrl    string          `json:"-"`
}

type OgpImageData struct {
	Url       string `json:"url"`
	SecureUrl string `json:"secure_url"`
	Type      string `json:"type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type OgpVideoData struct {
	Url       string `json:"url"`
	SecureUrl string `json:"secure_url"`
	Type      string `json:"type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type OgpAudioData struct {
	Url       string `json:"url"`
	SecureUrl string `json:"secure_url"`
	Type      string `json:"type"`
}

type OgpLocaleData struct {
	Locale    string `json:"locale"`
	Alternate string `json:"alternate"`
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

				ogp.Set(prop, cont)
			} else if n.Data == "title" {
				if len(ogp.Title) == 0 {
					ogp.Title = append(ogp.Title, n.FirstChild.Data)
				}
			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(root)
	if (len(ogp.Url) == 0) {
		ogp.Url = append(ogp.Url, url)
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

func (o *OgpData) Set(prop string, content string) {
	switch prop {
	case "og:title":
		o.Title = append(o.Title, content)
	case "og:type":
		o.Type = append(o.Type, content)
	case "og:url":
		o.Url = append(o.Url, content)
	case "og:image":
		fallthrough
	case "og:image:url":
		l := len(o.Image)
		if l > 0 && len(o.Image[l - 1].Url) == 0 {
			o.Image[l - 1].Url = content
		} else {
			o.Image = append(o.Image, OgpImageData{Url: content})
		}
	case "og:image:secure_url":
		l := len(o.Image)
		if l == 0 {
			o.Image = append(o.Image, OgpImageData{})
			l += 1
		}
		o.Image[l - 1].SecureUrl = content
	case "og:image:type":
		l := len(o.Image)
		if l == 0 {
			o.Image = append(o.Image, OgpImageData{})
			l += 1
		}
		o.Image[l - 1].Type = content
	case "og:image:width":
		l := len(o.Image)
		if l == 0 {
			o.Image = append(o.Image, OgpImageData{})
			l += 1
		}
		if num, err := strconv.Atoi(content); err != nil {
			console.Error("Failed to execute strconv.Atoi(): property=" + prop +
				", content=" + content + ", error=" + err.Error())
		} else {
			o.Image[l - 1].Width = num
		}
	case "og:image:height":
		l := len(o.Image)
		if l == 0 {
			o.Image = append(o.Image, OgpImageData{})
			l += 1
		}
		if num, err := strconv.Atoi(content); err != nil {
			console.Error("Failed to execute strconv.Atoi(): property=" + prop +
				", content=" + content + ", error=" + err.Error())
		} else {
			o.Image[l - 1].Height = num
		}
	case "og:video":
		fallthrough
	case "og:video:url":
		l := len(o.Video)
		if l > 0 && len(o.Video[l - 1].Url) == 0 {
			o.Video[l - 1].Url = content
		} else {
			o.Video = append(o.Video, OgpVideoData{Url: content})
		}
	case "og:video:secure_url":
		l := len(o.Video)
		if l == 0 {
			o.Video = append(o.Video, OgpVideoData{})
			l += 1
		}
		o.Video[l - 1].SecureUrl = content
	case "og:video:type":
		l := len(o.Video)
		if l == 0 {
			o.Video = append(o.Video, OgpVideoData{})
			l += 1
		}
		o.Video[l - 1].Type = content
	case "og:video:width":
		l := len(o.Video)
		if l == 0 {
			o.Video = append(o.Video, OgpVideoData{})
			l += 1
		}
		if num, err := strconv.Atoi(content); err != nil {
			console.Error("Failed to execute strconv.Atoi(): property=" + prop +
				", content=" + content + ", error=" + err.Error())
		} else {
			o.Video[l - 1].Width = num
		}
	case "og:video:height":
		l := len(o.Video)
		if l == 0 {
			o.Video = append(o.Video, OgpVideoData{})
			l += 1
		}
		if num, err := strconv.Atoi(content); err != nil {
			console.Error("Failed to execute strconv.Atoi(): property=" + prop +
				", content=" + content + ", error=" + err.Error())
		} else {
			o.Video[l - 1].Height = num
		}
	case "og:audio":
		fallthrough
	case "og:audio:url":
		l := len(o.Audio)
		if l > 0 && len(o.Video[l - 1].Url) == 0 {
			o.Audio[l - 1].Url = content
		} else {
			o.Audio = append(o.Audio, OgpAudioData{Url: content})
		}
	case "og:audio:secure_url":
		l := len(o.Audio)
		if l == 0 {
			o.Audio = append(o.Audio, OgpAudioData{})
			l += 1
		}
		o.Audio[l - 1].SecureUrl = content
	case "og:audio:type":
		l := len(o.Audio)
		if l == 0 {
			o.Audio = append(o.Audio, OgpAudioData{})
			l += 1
		}
		o.Audio[l - 1].Type = content
	case "og:description":
		o.Description = append(o.Description, content)
	case "og:determiner":
		o.Determiner = append(o.Determiner, content)
	case "og:locale":
		l := len(o.Locale)
		if l > 0 && len(o.Locale[l - 1].Locale) == 0 {
			o.Locale[l - 1].Locale = content
		} else {
			o.Locale = append(o.Locale, OgpLocaleData{Locale: content})
		}
	case "og:locale:alternate":
		l := len(o.Locale)
		if l == 0 {
			o.Locale = append(o.Locale, OgpLocaleData{})
			l += 1
		}
		o.Locale[l - 1].Alternate = content
	case "og:site_name":
		o.SiteName = append(o.SiteName, content)
	case "description":
		if len(o.Description) == 0 {
			o.Description = append(o.Description, content)
		}
	}
}