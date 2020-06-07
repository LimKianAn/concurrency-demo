package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	printNRandomUser(10)
}

func printNRandomUser(n int) {
	wg := &sync.WaitGroup{}
	mtx := &sync.RWMutex{}
	chCache := make(chan User)
	chDB := make(chan User)
	for i := 0; i < n; i++ {
		id := randomID()
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, mtx *sync.RWMutex, ch chan<- User) {
			if u, ok := queryCache(id, mtx); ok {
				ch <- u
			}

			wg.Done()
		}(id, wg, mtx, chCache)

		go func(id int, wg *sync.WaitGroup, mtx *sync.RWMutex, ch chan<- User) {
			if u, ok := queryDB(id); ok {
				mtx.Lock()
				cache[id] = u
				mtx.Unlock()
				ch <- u
			}

			wg.Done()
		}(id, wg, mtx, chDB)

		go func(chCache, chDB <-chan User) {
			select {
			case u := <-chCache:
				fmt.Println("Cache")
				fmt.Println(&u)

				<-chDB // clears the blocking msg in chDB
			case u := <-chDB:
				fmt.Println("DB")
				fmt.Println(&u)
			}
		}(chCache, chDB)

		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}
