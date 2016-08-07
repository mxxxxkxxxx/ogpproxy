package cache

import (
	"strings"
	"ogpproxy/ogpproxy"
)

type Writer interface {
	Write(data *ogpproxy.OgpData) error
}

type Reader interface {
	Read(url string) (*ogpproxy.OgpData, error)
}

type Handler interface {
	Writer
	Reader
}

func GetHandler() Handler {
	// @TODO: from config
	strategy := "LEVELDB"
	var handler Handler

	switch (strings.ToUpper(strategy)) {
	case "LEVELDB":
		handler = GetLevelDBHandler()
	// @TODO:
	// case "MEMCACHED":
	// 	handler = GetMemcachedHandler()
	default:
		panic("invalid storage handling strategy: " + strategy)
	}

	return handler
}
