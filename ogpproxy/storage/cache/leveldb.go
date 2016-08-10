package cache

import (
	"github.com/syndtr/goleveldb/leveldb"
	"encoding/json"
	"fmt"
	"ogpproxy/ogpproxy/ogp"
	"ogpproxy/ogpproxy/config"
	"ogpproxy/ogpproxy/console"
)

type LevelDBHandler struct {
	Path string
}

func (h *LevelDBHandler) Write(data *ogp.OgpData) error {
	var err error

	db, err := leveldb.OpenFile(h.Path, nil)
	if err != nil {
		return fmt.Errorf("Failed to open leveldb: err=[%s]", err)
	}
	defer db.Close()

	buf, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Failed to convert data to json: err=[%s]", err)
	}

	err = db.Put([]byte(data.RequestedUrl), buf, nil)
	if err != nil {
		return fmt.Errorf("Failed to write data to leveldb: err=[%s]", err)
	}

	console.Debug("Succeeded to write data to leveldb: url=[" + data.RequestedUrl + "]")

	return nil
}

func (h *LevelDBHandler) Read(url string) (*ogp.OgpData, error) {
	var err error
	data := &ogp.OgpData{}

	db, err := leveldb.OpenFile(h.Path, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to open leveldb: err=[%s]", err)
	}
	defer db.Close()

	buf, err := db.Get([]byte(url), nil)
	if err != nil || len(buf) == 0 {
		return nil, fmt.Errorf("Failed to get data from leveldb: url=[%s], err=[%s]", url, err)
	}

	// @TODO: check expiration
	err = json.Unmarshal(buf, data)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert data from json: url=[%s], err=[%s]", url, err)
	}

	return data, nil
}

func GetLevelDBHandler() *LevelDBHandler {
	handler := &LevelDBHandler{Path: config.GetConfig().LevelDB.Path}

	return handler
}