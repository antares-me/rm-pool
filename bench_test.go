package mypool

import (
	"sync"
	"testing"
)

type Counter struct{
	A int
	B int
}

func IncrementCounter (c *Counter) {
	c.A++
	c.B++
}

var counterPool = sync.Pool{
	New: func() interface{} { return new(Counter) },
}

func BenchmarkWithoutPoolWithNewStruct (b *testing.B) {
	var c *Counter
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			c = new(Counter)
			IncrementCounter(c)
		}
	}
}

func BenchmarkWithoutPoolWithoutNewStruct (b *testing.B) {
	c := new(Counter)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			IncrementCounter(c)
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var c *Counter
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			c = counterPool.Get().(*Counter)
			IncrementCounter(c)
			counterPool.Put(c)
		}
	}
}
