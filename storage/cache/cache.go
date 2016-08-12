package cache

import (
	"strings"
	"github.com/mxxxxkxxxx/ogpproxy/config"
)

type Writer interface {
	Write(key string, data []byte) error
	Remove(key string) error
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
	switch (strings.ToUpper(config.Cache.Strategy)) {
	case "LEVELDB":
		handler = GetLevelDBHandler()
	// @TODO:
	// case "MEMCACHED":
	// 	handler = GetMemcachedHandler()
	default:
		panic("invalid storage handling strategy: " + config.Cache.Strategy)
	}

	return handler
}
