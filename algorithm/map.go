package algorithm

import (
	"sync"
	"time"
)
/*
实现阻塞读且并发安全的map
思路
1、阻塞使用chan
2、并发安全使用sync.RWMutex
 */

type MyMap interface {
	Set(key string, val interface{})                   //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Get(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Key struct {
	c       chan bool // 用于唤醒获取该key阻塞的goroutine
	value   interface{}
	isExist bool // true表示已存在，false表示未存在，有goroutine阻塞
}

type Map struct {
	m   map[string]*Key
	RWm sync.RWMutex
}

func NewMap() *Map {
	return &Map{
		m: make(map[string]*Key),
		RWm: sync.RWMutex{},
	}
}

func (m *Map) Set(key string, val interface{}) {
	m.RWm.Lock()
	defer m.RWm.Unlock()

	v, ok := m.m[key]
	if !ok || v.isExist { // 该key不存在于map中
		tempKey := Key{
			value:   val,
			isExist: true,
		}
		m.m[key] = &tempKey
		return
	}

	if !v.isExist { // 该key已经存在，有goroutine阻塞
		v.isExist = true
		v.value = val
		close(v.c)
	}
}

func (m *Map) Get(key string, t time.Duration) interface{} {
	m.RWm.RLock()
	defer m.RWm.RUnlock()

	val, ok := m.m[key]
	if ok && val.isExist {
		return val.value
	}

	if !ok { // 不存在该key，阻塞或超时
		tempKey := Key{
			c:     make(chan bool),
			value: val,
		}
		m.m[key] = &tempKey

		go func(t time.Duration) {
			time.Sleep(t)
			close(tempKey.c)
		}(t)

		_ = <-tempKey.c
		if tempKey.isExist {
			return val.value
		}
		return nil
	}

	return nil
}
