package context

import (
	"github.com/becosuke/tasks-api/domain/entity/common"
	entity "github.com/becosuke/tasks-api/domain/entity/context"
	factoryCommon "github.com/becosuke/tasks-api/domain/factory/common"
	factory "github.com/becosuke/tasks-api/domain/factory/context"
	"github.com/becosuke/tasks-api/infrastructure/cache"
	database "github.com/becosuke/tasks-api/infrastructure/database/context"
	"github.com/becosuke/tasks-api/logger"
)

func FetchRecord(id uint64) (*entity.Record, error) {
	cacheKey := entity.GetRecordCacheKey(id)
	var res *entity.Record
	if cached, ok := cache.GetLocalCache(cacheKey, common.CacheExpireDefault); !ok {
		val, err := database.FindOne(id)
		if err != nil || !val.Valid() {
			return nil, err
		}

		cache.SetLocalCache(cacheKey, res)
		res = val
	} else {
		res = cached.(*entity.Record)
	}

	return res, nil
}

func Create(title string) (*entity.Document, error) {
	val, err := database.Create(title)
	if err != nil {
		return &entity.Document{}, err
	}

	return factory.Document(val), nil
}

func Update(id uint64, title string) (*entity.Document, error) {
	val, err := database.Update(id, title)
	if err != nil {
		return &entity.Document{}, err
	}

	cacheKey := entity.GetRecordCacheKey(id)
	cache.DeleteLocalCache(cacheKey)

	return factory.Document(val), nil
}

func Delete(id uint64) (*entity.Document, error) {
	val, err := database.Delete(id)
	if err != nil {
		return &entity.Document{}, err
	}

	cacheKey := entity.GetRecordCacheKey(id)
	cache.DeleteLocalCache(cacheKey)

	return factory.Document(val), nil
}

func FetchDocument(id uint64) (*entity.Document, error) {
	val, err := FetchRecord(id)
	if err != nil || !val.Valid() {
		return &entity.Document{}, err
	}

	return factory.Document(val), nil
}

func FetchDocuments(ids []uint64) (entity.Documents, error) {
	res := make(entity.Documents, 0, len(ids))
	for _, id := range ids {
		val, err := FetchDocument(id)
		if err != nil {
			return make(entity.Documents, 0), err
		}
		if val.Valid() {
			res = append(res, val)
		}
	}

	return res, nil
}

func FetchDocumentsAll(limit uint32, offset uint32) (entity.Documents, error) {
	cacheKey := entity.GetKeyAllCacheKey(limit, offset)
	var ids []uint64
	if cached, ok := cache.GetSharedCache(cacheKey); !ok {
		val, err := database.FindPrimaryKeyAll(limit, offset)
		if err != nil {
			return make(entity.Documents, 0), err
		}

		cache.SetSharedCache(cacheKey, val, common.CacheExpireDefault)
		ids = val
	} else {
		ids = cached.([]uint64)
		if ids == nil {
			logger.Print("ids is nil")
			ids = make([]uint64, 0)
		}
	}

	return FetchDocuments(ids)
}

func CountAll() (uint64, error) {
	cacheKey := entity.GetCountAllCacheKey()
	var res uint64
	if cached, ok := cache.GetSharedCache(cacheKey); !ok {
		val, err := database.CountAll()
		if err != nil {
			return 0, err
		}

		cache.SetSharedCache(cacheKey, val, common.CacheExpireDefault)
		res = val
	} else {
		res = cached.(uint64)
	}

	return res, nil
}

func FetchCountAll() (*common.Count, error) {
	count, err := CountAll()
	if err != nil {
		return &common.Count{}, err
	}

	return factoryCommon.Count(count), nil
}
