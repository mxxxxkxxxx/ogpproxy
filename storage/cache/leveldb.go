package cache

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/pkg/errors"
	"github.com/mxxxxkxxxx/ogpproxy/config"
	"github.com/mxxxxkxxxx/ogpproxy/console"
)

type LevelDBHandler struct {
	Path string
}

func (h *LevelDBHandler) Write(key string, data []byte) error {
	var err error

	if len(key) == 0 {
		return errors.New("Key is empty")
	}

	db, err := leveldb.OpenFile(h.Path, nil)
	if err != nil {
		return errors.Wrap(err, "Failed to open leveldb")
	}
	defer db.Close()

	err = db.Put([]byte(key), data, nil)
	if err != nil {
		return errors.Wrap(err, "Failed to write data to leveldb")
	}

	console.Debug("Succeeded to write data to leveldb: key=[" + key + "]")

	return nil
}

func (h *LevelDBHandler) Read(key string) ([]byte, error) {
	var err error

	if len(key) == 0 {
		return nil, errors.New("Key is empty")
	}

	db, err := leveldb.OpenFile(h.Path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open leveldb")
	}
	defer db.Close()

	buf, err := db.Get([]byte(key), nil)
	if err != nil || len(buf) == 0 {
		return nil, errors.Wrapf(err, "Failed to get data from leveldb: key=[%s]", key)
	}

	return buf, nil
}

func (h *LevelDBHandler) Remove(key string) error {
	var err error

	if len(key) == 0 {
		return errors.New("Key is empty")
	}

	db, err := leveldb.OpenFile(h.Path, nil)
	if err != nil {
		return errors.Wrap(err, "Failed to open leveldb")
	}
	defer db.Close()

	err = db.Delete([]byte(key), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to delete data from leveldb")
	}

	console.Debug("Succeeded to delete data from leveldb: key=[" + key + "]")

	return nil
}

func GetLevelDBHandler() *LevelDBHandler {
	handler := &LevelDBHandler{Path: config.GetConfig().LevelDB.Path}

	return handler
}