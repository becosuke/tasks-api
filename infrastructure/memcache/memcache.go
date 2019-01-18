package memcache

import (
	"github.com/becosuke/tasks-api/config"
	"github.com/bradfitz/gomemcache/memcache"
	"sync"
)

type connection struct {
	client *memcache.Client
	once   sync.Once
}

var conn = connection{}

func Open() connection {
	conn.once.Do(func() {
		conf := config.GetConfig()
		client := memcache.New(conf.Memcache.URL...)
		conn.client = client
	})

	return conn
}

func (c connection) Get(k string) ([]byte, error) {
	var value []byte
	var err error

	var v *memcache.Item
	if v, err = c.client.Get(k); err != nil {
		return value, err
	}
	value = v.Value
	return value, err
}

func (c connection) Set(k string, v []byte, e int32) error {
	return c.client.Set(&memcache.Item{
		Key:        k,
		Value:      v,
		Expiration: e,
	})
}

func (c connection) Delete(k string) error {
	return c.client.Delete(k)
}
