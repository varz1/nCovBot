package model

import (
	"github.com/varz1/nCovBot/cache"
	"sync"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := cache.New()
	data1 := ProvinceData{ProvinceName: "1"}
	data2 := ProvinceData{ProvinceName: "2"}
	c.Set("1", data1)
	d, i := c.Get("1")
	if !i{
		panic("")
	}
	data2 = d.(ProvinceData)
	t.Log(data2)
}

func TestTicker(t *testing.T) {
	t1 := time.NewTicker(1 * time.Second)
	defer t1.Stop()
	for {
		select {
		case <-t1.C:
			t.Log(time.Now())
		}
	}
}

func TestLock(t *testing.T) {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}
	t.Log(counter.m)
}
