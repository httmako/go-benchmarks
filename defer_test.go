package main

import (
	"sync"
	"testing"
)

func BenchmarkMutexDeferUnlock(b *testing.B) {
	var mu sync.Mutex
	for b.Loop() {
		func() {
			defer mu.Unlock()
			mu.Lock()
		}()
	}
}

func BenchmarkMutexUnlock(b *testing.B) {
	var mu sync.Mutex
	for b.Loop() {
		func() {
			mu.Lock()
			mu.Unlock()
		}()
	}
}
