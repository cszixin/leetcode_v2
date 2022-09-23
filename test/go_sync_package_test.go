package test

import (
	"sync"
	"testing"
)

var cs = 0
var mu sync.Mutex
var c = make(chan struct{}, 1)

func SyncByMutex() {
	mu.Lock()
	defer mu.Unlock()
	cs++
}

func SyncByChan() {
	c <- struct{}{}
	cs++
	<-c
}

func BenchmarkSyncByMutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SyncByMutex()
	}
}

func BenchmarkSyncByChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SyncByChan()
	}
}
