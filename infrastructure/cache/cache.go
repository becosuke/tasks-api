package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/becosuke/tasks-api/config"
	"github.com/becosuke/tasks-api/domain/entity/common"
	"github.com/becosuke/tasks-api/infrastructure/memcache"
	"sync"
)

var dictCachedAt = sync.Map{}
var dictLocalCache = sync.Map{}
var dictSharedCache = sync.Map{}

func getCachedAt(key string, expire int32) int64 {
	cacheKey := fmt.Sprintf("%s_cached_at", key)
	conf := config.GetConfig()

	var cachedAt int64
	localCachedAt, ok := dictCachedAt.Load(cacheKey)
	if ok == false || conf.NowTimestamp-localCachedAt.(int64) > common.CACHE_EXPIRE_LOCAL {
		mc := memcache.Open()
		if sharedCachedAt, mcError := mc.Get(cacheKey); mcError != nil {
			cachedAt = conf.NowTimestamp
			mc.Set(cacheKey, common.Int64ToBytes(cachedAt), expire)
		} else {
			cachedAt = common.BytesToInt64(sharedCachedAt)
		}
		dictCachedAt.Store(cacheKey, cachedAt)
	} else {
		cachedAt = localCachedAt.(int64)
	}

	return cachedAt
}

func GetLocalCache(key string, expire int32) (interface{}, bool) {
	cachedAt := getCachedAt(key, expire)

	entry, ok := dictLocalCache.Load(key)
	if ok == false || cachedAt > entry.(common.Cache).CachedAt {
		return nil, false
	} else {
		return entry.(common.Cache).CachedData, true
	}
}

func SetLocalCache(key string, data interface{}) {
	conf := config.GetConfig()

	entry := common.Cache{
		CachedData: data,
		CachedAt:   conf.NowTimestamp,
	}
	dictLocalCache.Store(key, entry)
}

func GetSharedCache(key string) (interface{}, bool) {
	conf := config.GetConfig()

	var data interface{}
	entry, ok := dictSharedCache.Load(key)
	if ok == false || conf.NowTimestamp-entry.(common.Cache).CachedAt > common.CACHE_EXPIRE_LOCAL {
		mc := memcache.Open()
		var tmp []byte
		var mcError, decError error
		tmp, mcError = mc.Get(key)
		if mcError != nil {
			return nil, false
		}
		buf := bytes.NewBuffer(tmp)
		decError = gob.NewDecoder(buf).Decode(&data)
		if decError != nil {
			return nil, false
		}

		entry = common.Cache{
			CachedData: data,
			CachedAt:   conf.NowTimestamp,
		}
		dictSharedCache.Store(key, entry)

		return data, true
	} else {
		return entry.(common.Cache).CachedData, true
	}
}

func GetSharedCache2(key string) (interface{}, bool) {
	conf := config.GetConfig()

	var data interface{}
	entry, ok := dictSharedCache.Load(key)
	if ok == false || conf.NowTimestamp-entry.(common.Cache).CachedAt > common.CACHE_EXPIRE_LOCAL {
		mc := memcache.Open()
		var tmp []byte
		var mcError, decError error
		tmp, mcError = mc.Get(key)
		if mcError != nil {
			return nil, false
		}
		buf := bytes.NewBuffer(tmp)
		decError = gob.NewDecoder(buf).Decode(&data)
		if decError != nil {
			return nil, false
		}

		entry = common.Cache{
			CachedData: data,
			CachedAt:   conf.NowTimestamp,
		}
		dictSharedCache.Store(key, entry)

		return data, true
	} else {
		return entry.(common.Cache).CachedData, true
	}
}

func SetSharedCache(key string, data interface{}, expire int32) {
	mc := memcache.Open()
	buf := bytes.NewBuffer(nil)
	if gobError := gob.NewEncoder(buf).Encode(&data); gobError == nil {
		mc.Set(key, buf.Bytes(), expire)
	}
}

func GetSharedNumber(key string) (uint64, bool) {
	var data uint64
	mc := memcache.Open()
	tmp, mcError := mc.Get(key)
	if mcError != nil {
		return 0, false
	}
	data = common.BytesToUint64(tmp)

	return data, true
}

func SetSharedNumber(key string, data uint64, expire int32) {
	mc := memcache.Open()
	tmp := common.Uint64ToBytes(data)
	mc.Set(key, tmp, expire)
}
