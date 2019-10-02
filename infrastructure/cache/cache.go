package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/becosuke/tasks-api/config"
	"github.com/becosuke/tasks-api/domain/entity/common"
	"github.com/becosuke/tasks-api/infrastructure/memcache"
)

var dictCachedAt = sync.Map{}
var dictLocalCache = sync.Map{}
var dictSharedCache = sync.Map{}

func getCachedAt(key string, expire int32) int64 {
	cacheKey := fmt.Sprintf("%s_cached_at", key)

	var cachedAt int64
	localCachedAt, ok := dictCachedAt.Load(cacheKey)
	if !ok || config.NowTimestamp()-localCachedAt.(int64) > common.CacheExpireLocal {
		mc := memcache.Open()
		if sharedCachedAt, mcError := mc.Get(cacheKey); mcError != nil {
			cachedAt = config.NowTimestamp()
			_ = mc.Set(cacheKey, common.Int64ToBytes(cachedAt), expire)
		} else {
			cachedAt = common.BytesToInt64(sharedCachedAt)
		}
		dictCachedAt.Store(cacheKey, cachedAt)
	} else {
		cachedAt = localCachedAt.(int64)
	}

	return cachedAt
}

func deleteCachedAt(key string) {
	cacheKey := fmt.Sprintf("%s_cached_at", key)

	dictLocalCache.Delete(cacheKey)

	mc := memcache.Open()
	_ = mc.Delete(cacheKey)
}

func GetLocalCache(key string, expire int32) (interface{}, bool) {
	cachedAt := getCachedAt(key, expire)

	entry, ok := dictLocalCache.Load(key)
	if !ok || cachedAt > entry.(common.Cache).CachedAt {
		return nil, false
	}

	return entry.(common.Cache).CachedData, true
}

func SetLocalCache(key string, data interface{}) {
	entry := common.Cache{
		CachedData: data,
		CachedAt:   config.NowTimestamp(),
	}
	dictLocalCache.Store(key, entry)
}

func DeleteLocalCache(key string) {
	dictLocalCache.Delete(key)
	deleteCachedAt(key)
}

func GetSharedCache(key string) (interface{}, bool) {
	var data interface{}
	entry, ok := dictSharedCache.Load(key)
	if !ok || config.NowTimestamp()-entry.(common.Cache).CachedAt > common.CacheExpireLocal {
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
			CachedAt:   config.NowTimestamp(),
		}
		dictSharedCache.Store(key, entry)

		return data, true
	}

	return entry.(common.Cache).CachedData, true
}

func SetSharedCache(key string, data interface{}, expire int32) {
	mc := memcache.Open()
	buf := bytes.NewBuffer(nil)
	if gobError := gob.NewEncoder(buf).Encode(&data); gobError == nil {
		_ = mc.Set(key, buf.Bytes(), expire)
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
	_ = mc.Set(key, tmp, expire)
}
