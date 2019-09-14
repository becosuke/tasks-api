package task

import (
	"github.com/becosuke/tasks-api/domain/entity/common"
	entity "github.com/becosuke/tasks-api/domain/entity/task"
	factoryCommon "github.com/becosuke/tasks-api/domain/factory/common"
	factory "github.com/becosuke/tasks-api/domain/factory/task"
	"github.com/becosuke/tasks-api/infrastructure/cache"
	database "github.com/becosuke/tasks-api/infrastructure/database/task"
)

func FetchRecord(id uint64) (*entity.Record, error) {
	var err error

	var res *entity.Record
	cacheKey := entity.GetRecordCacheKey(id)
	if cached, ok := cache.GetLocalCache(cacheKey, common.CacheExpireDefault); ok == false {
		if res, err = database.FindOne(id); err != nil || res.Valid() == false {
			return nil, err
		}

		cache.SetLocalCache(cacheKey, res)
	} else {
		res = cached.(*entity.Record)
	}

	return res, nil
}

func Create(listId uint64, title string) (*entity.Document, error) {
	var err error

	var val *entity.Record
	if val, err = database.Create(listId, title); err != nil {
		return &entity.Document{}, err
	}

	res := factory.Document(val)

	return res, nil
}

func Update(id uint64, title string) (*entity.Document, error) {
	var err error

	var val *entity.Record
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

	var val *entity.Record
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

	var val *entity.Record
	var res *entity.Document
	cacheKey := entity.GetDocumentCacheKey(id)
	if cached, ok := cache.GetLocalCache(cacheKey, common.CacheExpireDefault); ok == false {
		if val, err = FetchRecord(id); err != nil || val.Valid() == false {
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

func FetchDocumentsAll(limit uint32, offset uint32) ([]*entity.Document, error) {
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

func FetchDocumentsByRelationalKey(listId uint64, limit uint32, offset uint32) ([]*entity.Document, error) {
	var err error

	ids := make([]uint64, 0)
	cacheKey := entity.GetKeyRelationalCacheKey(listId, limit, offset)
	if cached, ok := cache.GetSharedCache(cacheKey); ok == false {
		if ids, err = database.FindPrimaryKeyByRelationalKey(listId, limit, offset); err != nil {
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

func CountByRelationalKey(listId uint64) (uint64, error) {
	cacheKey := entity.GetCountRelationalCacheKey(listId)
	var res uint64
	if cached, ok := cache.GetSharedCache(cacheKey); !ok {
		val, err := database.CountByRelationalKey(listId)
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

func FetchCountByRelationalKey(listId uint64) (*common.Count, error) {
	count, err := CountByRelationalKey(listId)
	if err != nil {
		return &common.Count{}, err
	}

	return factoryCommon.Count(count), nil
}

func FetchDocumentsByRelationalKeyAndProperties(listId uint64, contextIds []uint64, limit uint32, offset uint32) ([]*entity.Document, error) {
	var err error

	ids := make([]uint64, 0)
	cacheKey := entity.GetKeyRelationalAndPropertiesCacheKey(listId, contextIds, limit, offset)
	if cached, ok := cache.GetSharedCache(cacheKey); ok == false {
		if ids, err = database.FindPrimaryKeyByRelationalKeyAndProperties(listId, contextIds, limit, offset); err != nil {
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

func CountByRelationalKeyAndProperties(listId uint64, contextIds []uint64) (uint64, error) {
	cacheKey := entity.GetCountRelationalAndPropertiesCacheKey(listId, contextIds)
	var res uint64
	if cached, ok := cache.GetSharedCache(cacheKey); !ok {
		val, err := database.CountByRelationalKeyAndProperties(listId, contextIds)
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

func FetchCountByRelationalKeyAndProperties(listId uint64, contextIds []uint64) (*common.Count, error) {
	count, err := CountByRelationalKeyAndProperties(listId, contextIds)
	if err != nil {
		return &common.Count{}, err
	}

	return factoryCommon.Count(count), nil
}
