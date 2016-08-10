package cache

import (
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
	"ogpproxy/ogpproxy/config"
	"ogpproxy/ogpproxy/console"
)

type LevelDBHandler struct {
	Path string
}

func (h *LevelDBHandler) Write(key string, data []byte) error {
	var err error

	db, err := leveldb.OpenFile(h.Path, nil)
	if err != nil {
		return fmt.Errorf("Failed to open leveldb: err=[%s]", err)
	}
	defer db.Close()

	err = db.Put([]byte(key), data, nil)
	if err != nil {
		return fmt.Errorf("Failed to write data to leveldb: err=[%s]", err)
	}

	console.Debug("Succeeded to write data to leveldb: key=[" + key + "]")

	return nil
}

func (h *LevelDBHandler) Read(key string) ([]byte, error) {
	var err error

	db, err := leveldb.OpenFile(h.Path, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to open leveldb: err=[%s]", err)
	}
	defer db.Close()

	buf, err := db.Get([]byte(key), nil)
	if err != nil || len(buf) == 0 {
		return nil, fmt.Errorf("Failed to get data from leveldb: key=[%s], err=[%s]", key, err)
	}

	return buf, nil
}

func GetLevelDBHandler() *LevelDBHandler {
	handler := &LevelDBHandler{Path: config.GetConfig().LevelDB.Path}

	return handler
}