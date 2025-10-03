package main

import (
	"sync"
	"testing"
)

var m = map[string]string{}
var sm = sync.Map{}
var lm = map[string]string{}
var lock = sync.RWMutex{}

func AddToMap(key string, val string) {
	lock.Lock()
	defer lock.Unlock()
	lm[key] = val
}
func GetFromMap(key string) string {
	lock.RLock()
	defer lock.RUnlock()
	return lm[key]
}

func BenchmarkCreationMap(b *testing.B) {
	for b.Loop() {
		_ = map[string]string{}
	}
}

func BenchmarkCreationSyncMap(b *testing.B) {
	for b.Loop() {
		_ = sync.Map{}
	}
}

func BenchmarkAddMap(b *testing.B) {
	for b.Loop() {
		m[RandStringBytes(3)] = RandStringBytes(3)
		m[RandStringBytes(12)] = RandStringBytes(12)
	}
}

func BenchmarkAddLockMap(b *testing.B) {
	for b.Loop() {
		AddToMap(RandStringBytes(3), RandStringBytes(3))
		AddToMap(RandStringBytes(12), RandStringBytes(12))
	}
}

func BenchmarkAddSyncMap(b *testing.B) {
	for b.Loop() {
		sm.Store(RandStringBytes(3), RandStringBytes(3))
		sm.Store(RandStringBytes(12), RandStringBytes(12))
	}
}

func BenchmarkGetMap(b *testing.B) {
	for b.Loop() {
		_, _ = m[RandStringBytes(3)]
		_, _ = m[RandStringBytes(12)]
	}
}

func BenchmarkGetLockMap(b *testing.B) {
	for b.Loop() {
		_ = GetFromMap(RandStringBytes(3))
		_ = GetFromMap(RandStringBytes(12))
	}
}

func BenchmarkGetSyncMap(b *testing.B) {
	for b.Loop() {
		_, _ = sm.Load(RandStringBytes(3))
		_, _ = sm.Load(RandStringBytes(12))
	}
}

func BenchmarkSnapshotMap(b *testing.B) {
	for b.Loop() {
		snap := map[string]string{}
		for k, v := range m {
			snap[k] = v
		}
	}
}

func BenchmarkSnapshotSyncMap(b *testing.B) {
	for b.Loop() {
		snap := map[string]string{}
		sm.Range(func(key, val interface{}) bool {
			snap[key.(string)] = val.(string)
			return true
		})
	}
}
