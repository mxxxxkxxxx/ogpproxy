package page

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"github.com/mxxxxkxxxx/ogpproxy/controller"
	"encoding/json"
	"github.com/mxxxxkxxxx/ogpproxy/storage/cache"
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

	url := "https://mxxxxkxxxx.github.io/ogpproxy/samples/01.html"
	cache.GetHandler().Remove(url)

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
	}

	var expectedInt int
	var expectedStr string
	var expectedLen int

	expectedStr = "ogp proxy sample website"
	expectedLen = len(res.Ogp.Title)
	if expectedLen != 1 {
		t.Errorf("Title count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Title[0] != expectedStr {
		t.Errorf("Title is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Title[0])
	}

	expectedStr = "website"
	expectedLen = len(res.Ogp.Type)
	if expectedLen != 1 {
		t.Errorf("Type count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Type[0] != expectedStr {
		t.Errorf("Type is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Type[0])
	}

	expectedStr = "http://mxxxxkxxxx.github.io/ogpproxy/samples/01.html"
	expectedLen = len(res.Ogp.Url)
	if expectedLen != 1 {
		t.Errorf("Url count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Url[0] != expectedStr {
		t.Errorf("Url is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Url[0])
	}

	expectedStr = "This is a sample website for ogpproxy."
	expectedLen = len(res.Ogp.Description)
	if expectedLen != 1 {
		t.Errorf("Description count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Description[0] != expectedStr {
		t.Errorf("Description is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Description[0])
	}

	expectedStr = "auto"
	expectedLen = len(res.Ogp.Determiner)
	if expectedLen != 1 {
		t.Errorf("Determiner count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Determiner[0] != expectedStr {
		t.Errorf("Determiner is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Determiner[0])
	}

	expectedStr = "ogpproxy sample 01"
	expectedLen = len(res.Ogp.SiteName)
	if expectedLen != 1 {
		t.Errorf("SiteName count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.SiteName[0] != expectedStr {
		t.Errorf("SiteName is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.SiteName[0])
	}

	expectedLen = len(res.Ogp.Locale)
	if expectedLen != 1 {
		t.Errorf("Locale count must be 1, but actually %d.", expectedLen)
	} else {
		expectedStr = "en_US"
		if res.Ogp.Locale[0].Locale != expectedStr {
			t.Errorf("Locale[0].Locale is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Locale[0].Locale)
		}

		expectedStr = ""
		if res.Ogp.Locale[0].Alternate != expectedStr {
			t.Errorf("Locale[0].Alternate is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Locale[0].Alternate)
		}
	}

	expectedLen = len(res.Ogp.Audio)
	if expectedLen != 1 {
		t.Errorf("Audio count must be 1, but actually %d.", expectedLen)
	} else {
		expectedStr = "http://mxxxxkxxxx.github.io/ogpproxy/samples/sounds/01.mp3"
		if res.Ogp.Audio[0].Url != expectedStr {
			t.Errorf("Audio[0].Url is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Audio[0].Url)
		}

		expectedStr = ""
		if res.Ogp.Audio[0].SecureUrl != expectedStr {
			t.Errorf("Audio[0].SecureUrl is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Audio[0].SecureUrl)
		}

		expectedStr = ""
		if res.Ogp.Audio[0].Type != expectedStr {
			t.Errorf("Audio[0].Type is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Audio[0].Type)
		}
	}

	expectedLen = len(res.Ogp.Image)
	if expectedLen != 1 {
		t.Errorf("Image count must be 1, but actually %d.", expectedLen)
	} else {
		expectedStr = "http://mxxxxkxxxx.github.io/ogpproxy/samples/images/01.jpg"
		if res.Ogp.Image[0].Url != expectedStr {
			t.Errorf("Image[0].Url is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Image[0].Url)
		}

		expectedStr = ""
		if res.Ogp.Image[0].SecureUrl != expectedStr {
			t.Errorf("Image[0].SecureUrl is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Image[0].SecureUrl)
		}

		expectedStr = ""
		if res.Ogp.Image[0].Type != expectedStr {
			t.Errorf("Image[0].Type is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Image[0].Type)
		}

		expectedInt = 0
		if res.Ogp.Image[0].Width != expectedInt {
			t.Errorf("Image[0].Width is not expected one: expected=%d, actual=%d",
				expectedInt, res.Ogp.Image[0].Width)
		}

		expectedInt = 0
		if res.Ogp.Image[0].Height != expectedInt {
			t.Errorf("Image[0].Height is not expected one: expected=%d, actual=%d",
				expectedInt, res.Ogp.Image[0].Height)
		}
	}

	expectedLen = len(res.Ogp.Video)
	if expectedLen != 1 {
		t.Errorf("Video count must be 1, but actually %d.", expectedLen)
	} else {
		expectedStr = "http://mxxxxkxxxx.github.io/ogpproxy/samples/videos/01.mp4"
		if res.Ogp.Video[0].Url != expectedStr {
			t.Errorf("Video[0].Url is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Video[0].Url)
		}

		expectedStr = ""
		if res.Ogp.Video[0].SecureUrl != expectedStr {
			t.Errorf("Video[0].SecureUrl is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Video[0].SecureUrl)
		}

		expectedStr = ""
		if res.Ogp.Video[0].Type != expectedStr {
			t.Errorf("Video[0].Type is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Video[0].Type)
		}

		expectedInt = 0
		if res.Ogp.Video[0].Width != expectedInt {
			t.Errorf("Video[0].Width is not expected one: expected=%d, actual=%d",
				expectedInt, res.Ogp.Video[0].Width)
		}

		expectedInt = 0
		if res.Ogp.Video[0].Height != expectedInt {
			t.Errorf("Video[0].Height is not expected one: expected=%d, actual=%d",
				expectedInt, res.Ogp.Video[0].Height)
		}
	}
}


func TestGet_WithUrl_Valid_Full(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Get))
	defer server.Close()

	url := "https://mxxxxkxxxx.github.io/ogpproxy/samples/02.html"
	cache.GetHandler().Remove(url)

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
	}

	var expectedInt int
	var expectedStr string
	var expectedLen int

	expectedStr = "ogp proxy sample website"
	expectedLen = len(res.Ogp.Title)
	if expectedLen != 1 {
		t.Errorf("Title count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Title[0] != expectedStr {
		t.Errorf("Title is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Title[0])
	}

	expectedStr = "website"
	expectedLen = len(res.Ogp.Type)
	if expectedLen != 1 {
		t.Errorf("Type count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Type[0] != expectedStr {
		t.Errorf("Type is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Type[0])
	}

	expectedStr = "http://mxxxxkxxxx.github.io/ogpproxy/samples/01.html"
	expectedLen = len(res.Ogp.Url)
	if expectedLen != 1 {
		t.Errorf("Url count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Url[0] != expectedStr {
		t.Errorf("Url is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Url[0])
	}

	expectedStr = "This is a sample website for ogpproxy."
	expectedLen = len(res.Ogp.Description)
	if expectedLen != 1 {
		t.Errorf("Description count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Description[0] != expectedStr {
		t.Errorf("Description is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Description[0])
	}

	expectedStr = "auto"
	expectedLen = len(res.Ogp.Determiner)
	if expectedLen != 1 {
		t.Errorf("Determiner count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.Determiner[0] != expectedStr {
		t.Errorf("Determiner is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.Determiner[0])
	}

	expectedStr = "ogpproxy sample 01"
	expectedLen = len(res.Ogp.SiteName)
	if expectedLen != 1 {
		t.Errorf("SiteName count must be 1, but actually %d.", expectedLen)
	} else if res.Ogp.SiteName[0] != expectedStr {
		t.Errorf("SiteName is not expected one: expected=%s, actual=%s", expectedStr, res.Ogp.SiteName[0])
	}

	expectedLen = len(res.Ogp.Locale)
	if expectedLen != 1 {
		t.Errorf("Locale count must be 1, but actually %d.", expectedLen)
	} else {
		expectedStr = "en_US"
		if res.Ogp.Locale[0].Locale != expectedStr {
			t.Errorf("Locale[0].Locale is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Locale[0].Locale)
		}

		expectedStr = "ja_JP"
		if res.Ogp.Locale[0].Alternate != expectedStr {
			t.Errorf("Locale[0].Alternate is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Locale[0].Alternate)
		}
	}

	expectedLen = len(res.Ogp.Audio)
	if expectedLen != 1 {
		t.Errorf("Audio count must be 1, but actually %d.", expectedLen)
	} else {
		expectedStr = "http://mxxxxkxxxx.github.io/ogpproxy/samples/sounds/01.mp3"
		if res.Ogp.Audio[0].Url != expectedStr {
			t.Errorf("Audio[0].Url is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Audio[0].Url)
		}

		expectedStr = "https://mxxxxkxxxx.github.io/ogpproxy/samples/sounds/01.mp3"
		if res.Ogp.Audio[0].SecureUrl != expectedStr {
			t.Errorf("Audio[0].SecureUrl is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Audio[0].SecureUrl)
		}

		expectedStr = "audio/mpeg"
		if res.Ogp.Audio[0].Type != expectedStr {
			t.Errorf("Audio[0].Type is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Audio[0].Type)
		}
	}

	expectedLen = len(res.Ogp.Image)
	if expectedLen != 1 {
		t.Errorf("Image count must be 1, but actually %d.", expectedLen)
	} else {
		expectedStr = "http://mxxxxkxxxx.github.io/ogpproxy/samples/images/01.jpg"
		if res.Ogp.Image[0].Url != expectedStr {
			t.Errorf("Image[0].Url is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Image[0].Url)
		}

		expectedStr = "https://mxxxxkxxxx.github.io/ogpproxy/samples/images/01.jpg"
		if res.Ogp.Image[0].SecureUrl != expectedStr {
			t.Errorf("Image[0].SecureUrl is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Image[0].SecureUrl)
		}

		expectedStr = "image/jpeg"
		if res.Ogp.Image[0].Type != expectedStr {
			t.Errorf("Image[0].Type is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Image[0].Type)
		}

		expectedInt = 400
		if res.Ogp.Image[0].Width != expectedInt {
			t.Errorf("Image[0].Width is not expected one: expected=%d, actual=%d",
				expectedInt, res.Ogp.Image[0].Width)
		}

		expectedInt = 300
		if res.Ogp.Image[0].Height != expectedInt {
			t.Errorf("Image[0].Height is not expected one: expected=%d, actual=%d",
				expectedInt, res.Ogp.Image[0].Height)
		}
	}

	expectedLen = len(res.Ogp.Video)
	if expectedLen != 1 {
		t.Errorf("Video count must be 1, but actually %d.", expectedLen)
	} else {
		expectedStr = "http://mxxxxkxxxx.github.io/ogpproxy/samples/videos/01.mp4"
		if res.Ogp.Video[0].Url != expectedStr {
			t.Errorf("Video[0].Url is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Video[0].Url)
		}

		expectedStr = "https://mxxxxkxxxx.github.io/ogpproxy/samples/videos/01.mp4"
		if res.Ogp.Video[0].SecureUrl != expectedStr {
			t.Errorf("Video[0].SecureUrl is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Video[0].SecureUrl)
		}

		expectedStr = "video/mp4"
		if res.Ogp.Video[0].Type != expectedStr {
			t.Errorf("Video[0].Type is not expected one: expected=%s, actual=%s",
				expectedStr, res.Ogp.Video[0].Type)
		}

		expectedInt = 400
		if res.Ogp.Video[0].Width != expectedInt {
			t.Errorf("Video[0].Width is not expected one: expected=%d, actual=%d",
				expectedInt, res.Ogp.Video[0].Width)
		}

		expectedInt = 300
		if res.Ogp.Video[0].Height != expectedInt {
			t.Errorf("Video[0].Height is not expected one: expected=%d, actual=%d",
				expectedInt, res.Ogp.Video[0].Height)
		}
	}
}

func TestGet_WithUrl_Invalid(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Get))
	defer server.Close()

	url    := "mxxxxkxxxx.github.io/ogpproxy/samples/01.html"
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