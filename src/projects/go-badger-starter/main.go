package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dgraph-io/badger/v3"
)

type StoreResponse struct {
	Value string
	Meta  byte
}

type KVStoreInterface interface {
	Set(key string, value interface{}) error
	Get(key string) (StoreResponse, error)
}

type KVStore struct {
	DB *badger.DB
}

func NewBadgerDb(opts badger.Options) (*badger.DB, error) {
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewKVStore(db *badger.DB) KVStoreInterface {
	return &KVStore{
		DB: db,
	}
}

const (
	StringType  = 0x01
	IntType     = 0x02
	FloatType   = 0x03
	BoolType    = 0x04
	UnknownType = 0x05
)

// returns the type of the value - for usage in the meta
func getInterfaceType(value interface{}) byte {
	switch value.(type) {
	case int:
		return IntType
	case string:
		return StringType
	case float64:
		return FloatType
	case bool:
		return BoolType
	default:
		return UnknownType
	}
}

func (kv *KVStore) Set(key string, value interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return err
	}

	err = kv.DB.Update(func(txn *badger.Txn) error {
		valueType := getInterfaceType(value)
		entry := badger.NewEntry([]byte(key), buf.Bytes())
		entry = entry.WithMeta(valueType)
		err = txn.SetEntry(entry)
		return err
	})

	return err
}

func (kv *KVStore) Get(key string) (StoreResponse, error) {
	var resp StoreResponse
	err := kv.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

		if err != nil {
			return err
		}
		resp.Meta = item.UserMeta()
		err = item.Value(func(val []byte) error {
			dec := gob.NewDecoder(bytes.NewBuffer(val))
			if err := dec.Decode(&resp.Value); err != nil {
				return err
			}
			return nil
		})

		return err
	})

	if err != nil {
		return StoreResponse{}, err
	}

	return resp, nil
}

func main() {
	// open the badger database
	// define the badger database options
	badgerOpts := badger.DefaultOptions("./badgerdb")
	badgerOpts.ValueDir = "./badgerdb"
	// deactivate badger logger
	badgerOpts.Logger = nil

	db, err := NewBadgerDb(badgerOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := NewKVStore(db)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/store/:key", func(c echo.Context) error {
		key := c.Param("key")
		if c.Request().Body == nil {
			return c.String(http.StatusBadRequest, "Invalid or Empty data.")
		}
		body, _ := ioutil.ReadAll(c.Request().Body)
		value := string(body)
		err := store.Set(key, value)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusCreated, "Successful!")
	})

	e.GET("/store/:key", func(c echo.Context) error {
		key := c.Param("key")

		resp, err := store.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return c.String(http.StatusNotFound, "Not Found")
			}
			return c.String(500, err.Error())
		}
		value := resp.Value
		return c.String(http.StatusOK, value)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
