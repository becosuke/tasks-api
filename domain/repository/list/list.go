package list

import (
	"github.com/becosuke/tasks-api/domain/entity/common"
	entity "github.com/becosuke/tasks-api/domain/entity/list"
	factory "github.com/becosuke/tasks-api/domain/factory/list"
	"github.com/becosuke/tasks-api/infrastructure/cache"
	database "github.com/becosuke/tasks-api/infrastructure/database/list"
)

func FetchEntity(id uint64) (*entity.Entity, error) {
	var err error

	var res *entity.Entity
	cacheKey := entity.GetEntityCacheKey(id)
	if cached, ok := cache.GetLocalCache(cacheKey, common.CacheExpireDefault); ok == false {
		if res, err = database.FindOne(id); err != nil || res.Valid() == false {
			return nil, err
		}

		cache.SetLocalCache(cacheKey, res)
	} else {
		res = cached.(*entity.Entity)
	}

	return res, nil
}

func Create(title string) (*entity.Document, error) {
	var err error

	var val *entity.Entity
	if val, err = database.Create(title); err != nil {
		return &entity.Document{}, err
	}

	res := factory.Document(val)

	return res, nil
}

func Update(id uint64, title string) (*entity.Document, error) {
	var err error

	var val *entity.Entity
	if val, err = database.Update(id, title); err != nil {
		return &entity.Document{}, err
	}

	cacheKey := entity.GetDocumentCacheKey(id)
	cache.DeleteLocalCache(cacheKey)

	res := factory.Document(val)

	return res, nil
}

func Delete(id uint64) (*entity.Document, error) {
	var err error

	var val *entity.Entity
	if val, err = database.Delete(id); err != nil {
		return &entity.Document{}, err
	}

	cacheKey := entity.GetDocumentCacheKey(id)
	cache.DeleteLocalCache(cacheKey)

	res := factory.Document(val)

	return res, nil
}

func FetchDocument(id uint64) (*entity.Document, error) {
	var err error

	var val *entity.Entity
	var res *entity.Document
	cacheKey := entity.GetDocumentCacheKey(id)
	if cached, ok := cache.GetLocalCache(cacheKey, common.CacheExpireDefault); ok == false {
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

func FetchDocumentsAll(limit int32, offset int32) ([]*entity.Document, error) {
	var err error

	ids := make([]uint64, 0)
	cacheKey := entity.GetKeyAllCacheKey(limit, offset)
	if cached, ok := cache.GetSharedCache(cacheKey); ok == false {
		if ids, err = database.FindPrimaryKeyAll(limit, offset); err != nil {
			return make([]*entity.Document, 0), err
		}

		cache.SetSharedCache(cacheKey, ids, common.CacheExpireDefault)
	} else {
		if ids = cached.([]uint64); ids == nil {
			ids = make([]uint64, 0)
		}
	}

	return FetchDocuments(ids)
}

func CountAll() (uint64, error) {
	var err error

	var res uint64
	cacheKey := entity.GetCountAllCacheKey()
	if cached, ok := cache.GetSharedCache(cacheKey); ok == false {
		if res, err = database.CountAll(); err != nil {
			return 0, err
		}

		cache.SetSharedCache(cacheKey, res, common.CacheExpireDefault)
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

	return common.NewCount(count), nil
}
