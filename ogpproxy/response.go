package ogpproxy

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type OgpData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Image       string `json:"image"`
	SiteName    string `json:"site_name"`
	Locale      string `json:"locale"`
}

type Response struct {
	Writer http.ResponseWriter `json:"-"`
	Errors []string            `json:"errors"`
	Ogp    *OgpData             `json:"ogp"`
}

func (res *Response) Write() {
	var content string
	buf, err := json.Marshal(res)
	if err == nil {
		content = string(buf)
	} else {
		fmt.Printf("error: %s", err)
		content = "{errors: [\"An unexpected error occured.\"]}"
	}

	res.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprintf(res.Writer, content)
}

func (res *Response) AddError(msg string) {
	res.Errors = append(res.Errors, msg)
}

func (res *Response) WriteError(msg string) {
	res.AddError(msg)
	res.Write()
}
