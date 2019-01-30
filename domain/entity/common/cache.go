package common

const CacheExpireOneMinute = 60
const CacheExpireHalfHour = 1800
const CacheExpireOneHour = 3600
const CacheExpireOneThirdDay = 28800
const CacheExpireHalfDay = 43200
const CacheExpireDefault = CacheExpireHalfHour
const CacheExpireLocal = CacheExpireOneMinute

type Cache struct {
	CachedData interface{}
	CachedAt   int64
}
