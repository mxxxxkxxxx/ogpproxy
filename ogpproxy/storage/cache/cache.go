package cache

import (
	"strings"
	"ogpproxy/ogpproxy/ogp"
	"ogpproxy/ogpproxy/config"
)

type Writer interface {
	Write(data *ogp.OgpData) error
}

type Reader interface {
	Read(url string) (*ogp.OgpData, error)
}

type Handler interface {
	Writer
	Reader
}

func GetHandler() Handler {
	var handler Handler

	config := config.GetConfig()
	switch (strings.ToUpper(config.Cache)) {
	case "LEVELDB":
		handler = GetLevelDBHandler()
	// @TODO:
	// case "MEMCACHED":
	// 	handler = GetMemcachedHandler()
	default:
		panic("invalid storage handling strategy: " + config.Cache)
	}

	return handler
}
