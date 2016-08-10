package page

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"ogpproxy/ogpproxy/controller"
	"encoding/json"
)

func TestGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Get))
	defer server.Close()

	r, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Failed to do GET: err=[%+v]", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Failed to read body: err=[%+v]", err)
	}

	res := &controller.Response{}
	json.Unmarshal(data, res)
	if len(res.Errors) != 1 {
		t.Fatalf("Error count must equal 1: actual=%d", len(res.Errors))
	} else if res.Errors[0] != "You must specify a variable named url." {
		t.Fatalf("Error message is not expected one: actual=%s", res.Errors[0])
	}
}

func TestGet_WithUrl_Valid(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Get))
	defer server.Close()

	url    := "https://www.google.com/"
	r, err := http.Get(server.URL + "?url=" + url)
	if err != nil {
		t.Fatalf("Failed to do GET: err=[%+v]", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Failed to read body: err=[%+v]", err)
	}

	res := &controller.Response{}
	json.Unmarshal(data, res)
	if len(res.Errors) != 0 {
		t.Fatalf("Error count must equal 0: actual=%+v", res.Errors)
	} else if res.Ogp.Title != "Google" {
		t.Fatalf("Title is not expected one: actual=%s", res.Ogp.Title)
	}
}

func TestGet_WithUrl_Invalid(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Get))
	defer server.Close()

	url    := "www.google.com"
	r, err := http.Get(server.URL + "?url=" + url)
	if err != nil {
		t.Fatalf("Failed to do GET: err=[%+v]", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Failed to read body: err=[%+v]", err)
	}

	res := &controller.Response{}
	json.Unmarshal(data, res)
	if len(res.Errors) != 1 {
		t.Fatalf("Error count must equal 1: actual=%d", len(res.Errors))
	} else if res.Errors[0] != "Failed to do GET " + url {
		t.Fatalf("Error message is not expected one: actual=%s", res.Errors[0])
	}
}

func TestGet_WithUrl_NotExist(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Get))
	defer server.Close()

	url    := "http://this.must.not.exist.com/"
	r, err := http.Get(server.URL + "?url=" + url)
	if err != nil {
		t.Fatalf("Failed to do GET: err=[%+v]", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Failed to read body: err=[%+v]", err)
	}

	res := &controller.Response{}
	json.Unmarshal(data, res)
	if len(res.Errors) != 1 {
		t.Fatalf("Error count must equal 1: actual=%d", len(res.Errors))
	} else if res.Errors[0] != "Failed to do GET " + url {
		t.Fatalf("Error message is not expected one: actual=%s", res.Errors[0])
	}
}