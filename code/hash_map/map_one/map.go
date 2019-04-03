package map_one

import "sync"

var SHARD_COUNT = 32

type ConCurrentMap []*ElementMap

type ElementMap struct {
	m map[string]interface{}
	sync.RWMutex
}

func New() ConCurrentMap {
	m := make([]*ElementMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++{
		m[i] = &ElementMap{m: map[string]interface{}{}}
	}

	return m
}

// 获取分片元素
func (m ConCurrentMap) GetShare(k string) *ElementMap{
	return m[uint(fnv32(k))%uint(SHARD_COUNT)]
}

func (m ConCurrentMap) MSet(data map[string]interface{})  {
	for k, v := range data{
		shard := m.GetShare(k)
		shard.Lock()
		shard.m[k] = v
		shard.Unlock()
	}
}

// FNV hash算法
func fnv32(k string) uint32  {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(k); i++ {
		hash *= prime32
		hash ^= uint32(k[i])
	}
	return hash
}