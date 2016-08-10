package controller

import (
	"encoding/json"
	"net/http"
	"fmt"
	"ogpproxy/ogpproxy/ogp"
	"ogpproxy/ogpproxy/console"
)

type Response struct {
	Writer http.ResponseWriter `json:"-"`
	Errors []string            `json:"errors"`
	Ogp    *ogp.OgpData        `json:"ogp"`
}

func (res *Response) Write() {
	var content string
	buf, err := json.Marshal(res)
	if err == nil {
		content = string(buf)
	} else {
		console.Error("Failed to write response: err=[" + err.Error() + "]")
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
