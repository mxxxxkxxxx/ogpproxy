package cache

import (
	"ogpproxy/ogpproxy"
	"github.com/syndtr/goleveldb/leveldb"
	"encoding/json"
	"fmt"
)

type LevelDBHandler struct {}

func (h *LevelDBHandler) Write(data *ogpproxy.OgpData) error {
	var err error

	db, err := leveldb.OpenFile("cache.db", nil)
	if err != nil {
		return fmt.Errorf("Failed to open leveldb: err=[%s]", err)
	}
	defer db.Close()

	buf, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Failed to convert data to json: err=[%s]", err)
	}

	err = db.Put([]byte(data.Url), buf, nil)
	if err != nil {
		return fmt.Errorf("Failed to write data to leveldb: err=[%s]", err)
	}

	fmt.Printf("Succeeded to write data to leveldb: url=[%s]\n", data.Url)

	return nil
}

func (h *LevelDBHandler) Read(url string) (*ogpproxy.OgpData, error) {
	var err error
	data := &ogpproxy.OgpData{}

	db, err := leveldb.OpenFile("cache.db", nil)
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
	handler := &LevelDBHandler{}

	return handler
}