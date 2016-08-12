package ogp

import (
	"golang.org/x/net/html"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/mxxxxkxxxx/ogpproxy/storage/cache"
	"strconv"
	"github.com/mxxxxkxxxx/ogpproxy/console"
	"time"
	"fmt"
	"github.com/mxxxxkxxxx/ogpproxy/config"
)

type OgpData struct {
	Core         *OgpDataCore `json:"core"`
	CreatedAt    time.Time   `json:"created_at"`
	RequestedUrl string      `json:"requested_url"`
}

type OgpDataCore struct {
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
	ogp := &OgpData{Core:&OgpDataCore{}}

	var f func(n *html.Node)
	var title, desc string
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "meta" {
				var prop, cont, name string
				for _, attr := range n.Attr {
					switch attr.Key {
					case "property":
						prop = attr.Val
					case "content":
						cont = attr.Val
					case "name":
						name = attr.Val
					}
				}

				if len(cont) == 0 {
					return
				} else if len(prop) == 0 {
					if len(name) > 0 && name == "description" {
						desc = cont
					}

					return
				}

				ogp.Set(prop, cont)
			} else if n.Data == "title" {
				if len(ogp.Core.Title) == 0 {
					title = n.FirstChild.Data
				}
			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(root)

	if (len(ogp.Core.Title) == 0 && len(title) > 0) {
		ogp.Core.Title = append(ogp.Core.Title, title)
	}

	if (len(ogp.Core.Description) == 0 && len(desc) > 0) {
		ogp.Core.Description = append(ogp.Core.Description, desc)
	}

	if (len(ogp.Core.Url) == 0) {
		ogp.Core.Url = append(ogp.Core.Url, url)
	}

	ogp.RequestedUrl = url
	ogp.CreatedAt    = time.Now()

	return ogp
}

func LoadOgpData(url string) (*OgpData, error) {
	data := &OgpData{}

	cacheHandler := cache.GetHandler()
	buf, err := cacheHandler.Read(url)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to load ogp data: key=[%s]", url)
	}

	err = json.Unmarshal(buf, data)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to convert ogp data from json: key=[%s]", url)
	}

	duration := int(time.Now().Sub(data.CreatedAt).Seconds())
	console.Debug(fmt.Sprintf("ogp data duration: %d", duration))

	if duration > config.GetConfig().Cache.Expiration {
		console.Debug("Remove a expired cache: key=[" + url + "]")
		err = Delete(data)
		if err != nil {
			console.Error("Failed to remove a expired cache: key=[" + url + "], err=[" + err.Error() + "]")
		} else {
			return nil, fmt.Errorf("Succeeded to remove a expired cache: key=[" + url + "]")
		}
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

func Delete(o *OgpData) error {
	cacheHandler := cache.GetHandler()
	err := cacheHandler.Remove(o.RequestedUrl)
	if err != nil {
		return errors.Wrapf(err, "Failed to delete ogp data: key=[%s]", o.RequestedUrl)
	}

	return nil
}

func (o *OgpData) Set(prop string, content string) {
	switch prop {
	case "og:title":
		o.Core.Title = append(o.Core.Title, content)
	case "og:type":
		o.Core.Type = append(o.Core.Type, content)
	case "og:url":
		o.Core.Url = append(o.Core.Url, content)
	case "og:image":
		fallthrough
	case "og:image:url":
		l := len(o.Core.Image)
		if l > 0 && len(o.Core.Image[l - 1].Url) == 0 {
			o.Core.Image[l - 1].Url = content
		} else {
			o.Core.Image = append(o.Core.Image, OgpImageData{Url: content})
		}
	case "og:image:secure_url":
		l := len(o.Core.Image)
		if l == 0 {
			o.Core.Image = append(o.Core.Image, OgpImageData{})
			l += 1
		}
		o.Core.Image[l - 1].SecureUrl = content
	case "og:image:type":
		l := len(o.Core.Image)
		if l == 0 {
			o.Core.Image = append(o.Core.Image, OgpImageData{})
			l += 1
		}
		o.Core.Image[l - 1].Type = content
	case "og:image:width":
		l := len(o.Core.Image)
		if l == 0 {
			o.Core.Image = append(o.Core.Image, OgpImageData{})
			l += 1
		}
		if num, err := strconv.Atoi(content); err != nil {
			console.Error("Failed to execute strconv.Atoi(): property=" + prop +
				", content=" + content + ", error=" + err.Error())
		} else {
			o.Core.Image[l - 1].Width = num
		}
	case "og:image:height":
		l := len(o.Core.Image)
		if l == 0 {
			o.Core.Image = append(o.Core.Image, OgpImageData{})
			l += 1
		}
		if num, err := strconv.Atoi(content); err != nil {
			console.Error("Failed to execute strconv.Atoi(): property=" + prop +
				", content=" + content + ", error=" + err.Error())
		} else {
			o.Core.Image[l - 1].Height = num
		}
	case "og:video":
		fallthrough
	case "og:video:url":
		l := len(o.Core.Video)
		if l > 0 && len(o.Core.Video[l - 1].Url) == 0 {
			o.Core.Video[l - 1].Url = content
		} else {
			o.Core.Video = append(o.Core.Video, OgpVideoData{Url: content})
		}
	case "og:video:secure_url":
		l := len(o.Core.Video)
		if l == 0 {
			o.Core.Video = append(o.Core.Video, OgpVideoData{})
			l += 1
		}
		o.Core.Video[l - 1].SecureUrl = content
	case "og:video:type":
		l := len(o.Core.Video)
		if l == 0 {
			o.Core.Video = append(o.Core.Video, OgpVideoData{})
			l += 1
		}
		o.Core.Video[l - 1].Type = content
	case "og:video:width":
		l := len(o.Core.Video)
		if l == 0 {
			o.Core.Video = append(o.Core.Video, OgpVideoData{})
			l += 1
		}
		if num, err := strconv.Atoi(content); err != nil {
			console.Error("Failed to execute strconv.Atoi(): property=" + prop +
				", content=" + content + ", error=" + err.Error())
		} else {
			o.Core.Video[l - 1].Width = num
		}
	case "og:video:height":
		l := len(o.Core.Video)
		if l == 0 {
			o.Core.Video = append(o.Core.Video, OgpVideoData{})
			l += 1
		}
		if num, err := strconv.Atoi(content); err != nil {
			console.Error("Failed to execute strconv.Atoi(): property=" + prop +
				", content=" + content + ", error=" + err.Error())
		} else {
			o.Core.Video[l - 1].Height = num
		}
	case "og:audio":
		fallthrough
	case "og:audio:url":
		l := len(o.Core.Audio)
		if l > 0 && len(o.Core.Audio[l - 1].Url) == 0 {
			o.Core.Audio[l - 1].Url = content
		} else {
			o.Core.Audio = append(o.Core.Audio, OgpAudioData{Url: content})
		}
	case "og:audio:secure_url":
		l := len(o.Core.Audio)
		if l == 0 {
			o.Core.Audio = append(o.Core.Audio, OgpAudioData{})
			l += 1
		}
		o.Core.Audio[l - 1].SecureUrl = content
	case "og:audio:type":
		l := len(o.Core.Audio)
		if l == 0 {
			o.Core.Audio = append(o.Core.Audio, OgpAudioData{})
			l += 1
		}
		o.Core.Audio[l - 1].Type = content
	case "og:description":
		o.Core.Description = append(o.Core.Description, content)
	case "og:determiner":
		o.Core.Determiner = append(o.Core.Determiner, content)
	case "og:locale":
		l := len(o.Core.Locale)
		if l > 0 && len(o.Core.Locale[l - 1].Locale) == 0 {
			o.Core.Locale[l - 1].Locale = content
		} else {
			o.Core.Locale = append(o.Core.Locale, OgpLocaleData{Locale: content})
		}
	case "og:locale:alternate":
		l := len(o.Core.Locale)
		if l == 0 {
			o.Core.Locale = append(o.Core.Locale, OgpLocaleData{})
			l += 1
		}
		o.Core.Locale[l - 1].Alternate = content
	case "og:site_name":
		o.Core.SiteName = append(o.Core.SiteName, content)
	}
}