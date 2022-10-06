/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-29 14:48:14
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-29 14:53:06
 */
package test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	sync.RWMutex
	data string
}

func BenchmarkAtomicSet(b *testing.B) {
	var config atomic.Value
	c := Config{data: "hello"}
	b.ReportAllocs()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			config.Store(c)
		}
	})
}

func BenchmarkAtomicGet(b *testing.B) {
	var config atomic.Value
	config.Store(Config{data: "hello"})
	b.ReportAllocs()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = config.Load().(Config)
		}
	})
}
