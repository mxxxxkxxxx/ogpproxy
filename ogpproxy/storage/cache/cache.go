package cache

import (
	"strings"
	"ogpproxy/ogpproxy/config"
)

type Writer interface {
	Write(key string, data []byte) error
}

type Reader interface {
	Read(key string) ([]byte, error)
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
