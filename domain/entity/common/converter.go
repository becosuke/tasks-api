package common

import (
	"strconv"
)

func Uint64ToBytes(n uint64) []byte {
	s := strconv.FormatUint(n, 10)
	return []byte(s)
}

func BytesToUint64(bs []byte) uint64 {
	s := string(bs)
	n, e := strconv.ParseUint(s, 10, 64)
	if e != nil {
		return 0
	}

	return n
}

func Int64ToBytes(n int64) []byte {
	s := strconv.FormatInt(n, 10)
	return []byte(s)
}

func BytesToInt64(bs []byte) int64 {
	s := string(bs)
	n, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		return 0
	}

	return n
}

func Extract(ns []uint64, limit int32, offset int32) []uint64 {
	if limit == 0 {
		return make([]uint64, 0)
	}

	if int32(len(ns)) <= offset {
		return make([]uint64, 0)
	}

	if int32(len(ns)) < offset+limit {
		return ns[offset:]
	}

	return ns[offset : offset+limit]
}
