package task

import (
	"github.com/becosuke/tasks-api/domain/entity/common"
	entity "github.com/becosuke/tasks-api/domain/entity/task"
	factory "github.com/becosuke/tasks-api/domain/factory/task"
	"github.com/becosuke/tasks-api/infrastructure/cache"
	database "github.com/becosuke/tasks-api/infrastructure/database/task"
)

func FetchEntity(id uint64) (*entity.Entity, error) {
	var err error

	var res *entity.Entity
	cacheKey := entity.GetEntityCacheKey(id)
	if cached, ok := cache.GetLocalCache(cacheKey, common.CACHE_EXPIRE_DEFAULT); ok == false {
		if res, err = database.FindOne(id); err != nil || res.Valid() == false {
			return nil, err
		}

		cache.SetLocalCache(cacheKey, res)
	} else {
		res = cached.(*entity.Entity)
	}

	return res, nil
}

func FetchDocument(id uint64) (*entity.Document, error) {
	var err error

	var val *entity.Entity
	var res *entity.Document
	cacheKey := entity.GetDocumentCacheKey(id)
	if cached, ok := cache.GetLocalCache(cacheKey, common.CACHE_EXPIRE_DEFAULT); ok == false {
		if val, err = FetchEntity(id); err != nil || val.Valid() == false {
			return &entity.Document{}, err
		}

		res = factory.Document(val)
		cache.SetLocalCache(cacheKey, res)
	} else {
		res = cached.(*entity.Document)
	}

	return res, nil
}

func FetchDocuments(ids []uint64) ([]*entity.Document, error) {
	res := make([]*entity.Document, 0, len(ids))
	for _, id := range ids {
		if val, err := FetchDocument(id); err != nil {
			return make([]*entity.Document, 0), err
		} else if val.Valid() {
			res = append(res, val)
		}
	}

	return res, nil
}

func CountAll() (uint64, error) {
	var err error

	var res uint64
	cacheKey := entity.GetCountAllCacheKey()
	if cached, ok := cache.GetSharedCache(cacheKey); ok == false {
		if res, err = database.CountAll(); err != nil {
			return 0, err
		}

		cache.SetSharedCache(cacheKey, res, common.CACHE_EXPIRE_DEFAULT)
	} else {
		res = cached.(uint64)
	}

	return res, nil
}

func FetchAll(limit int32, offset int32) ([]*entity.Document, error) {
	var err error

	ids := make([]uint64, 0)
	cacheKey := entity.GetKeyAllCacheKey(limit, offset)
	if cached, ok := cache.GetSharedCache(cacheKey); ok == false {
		if ids, err = database.FindPrimaryKeyAll(limit, offset); err != nil {
			return make([]*entity.Document, 0), err
		}

		cache.SetSharedCache(cacheKey, ids, common.CACHE_EXPIRE_DEFAULT)
	} else {
		if ids = cached.([]uint64); ids == nil {
			ids = make([]uint64, 0)
		}
	}

	return FetchDocuments(ids)
}