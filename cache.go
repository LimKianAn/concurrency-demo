package main

import "sync"

var cache = map[int]User{}

func queryCache(id int, mtx *sync.RWMutex) (User, bool) {
	mtx.RLock()
	u, ok := cache[id]
	mtx.RUnlock()

	return u, ok
}
