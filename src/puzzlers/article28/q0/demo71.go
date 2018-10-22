package main

import (
	"fmt"
	"sync"
)

// ConcurrentMap 代表自制的简易并发安全字典。
type ConcurrentMap struct {
	m  map[interface{}]interface{}
	mu sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[interface{}]interface{}),
	}
}

func (cMap *ConcurrentMap) Delete(key interface{}) {
	cMap.mu.Lock()
	defer cMap.mu.Unlock()
	delete(cMap.m, key)
}

func (cMap *ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	cMap.mu.RLock()
	defer cMap.mu.RUnlock()
	value, ok = cMap.m[key]
	return
}

func (cMap *ConcurrentMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	cMap.mu.Lock()
	defer cMap.mu.Unlock()
	actual, loaded = cMap.m[key]
	if loaded {
		return
	}
	cMap.m[key] = value
	actual = value
	return
}

func (cMap *ConcurrentMap) Range(f func(key, value interface{}) bool) {
	cMap.mu.RLock()
	defer cMap.mu.RUnlock()
	for k, v := range cMap.m {
		if !f(k, v) {
			break
		}
	}
}

func (cMap *ConcurrentMap) Store(key, value interface{}) {
	cMap.mu.Lock()
	defer cMap.mu.Unlock()
	cMap.m[key] = value
}

func main() {
	pairs := []struct {
		k int
		v string
	}{
		{k: 1, v: "a"},
		{k: 2, v: "b"},
		{k: 3, v: "c"},
		{k: 4, v: "d"},
	}

	// 示例1。
	{
		cMap := NewConcurrentMap()
		cMap.Store(pairs[0].k, pairs[0].v)
		cMap.Store(pairs[1].k, pairs[1].v)
		cMap.Store(pairs[2].k, pairs[2].v)
		fmt.Println("[Three pairs have been stored in the ConcurrentMap instance]")

		cMap.Range(func(key, value interface{}) bool {
			fmt.Printf("The result of an iteration in Range: %v, %v\n",
				key, value)
			return true
		})

		k0 := pairs[0].k
		v0, ok := cMap.Load(k0)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v0, ok, k0)

		k3 := pairs[3].k
		v3, ok := cMap.Load(k3)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v3, ok, k3)

		k2, v2 := pairs[2].k, pairs[2].v
		actual2, loaded2 := cMap.LoadOrStore(k2, v2)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual2, loaded2, k2, v2)
		v3 = pairs[3].v
		actual3, loaded3 := cMap.LoadOrStore(k3, v3)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual3, loaded3, k3, v3)

		k1 := pairs[1].k
		cMap.Delete(k1)
		fmt.Printf("[The pair with the key of %v has been removed from the ConcurrentMap instance]\n",
			k1)
		v1, ok := cMap.Load(k1)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v1, ok, k1)
		v1 = pairs[1].v
		actual1, loaded1 := cMap.LoadOrStore(k1, v1)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual1, loaded1, k1, v1)

		cMap.Range(func(key, value interface{}) bool {
			fmt.Printf("The result of an iteration in Range: %v, %v\n",
				key, value)
			return true
		})
	}
	fmt.Println()

	// 示例2。
	{
		var sMap sync.Map
		sMap.Store(pairs[0].k, pairs[0].v)
		sMap.Store(pairs[1].k, pairs[1].v)
		sMap.Store(pairs[2].k, pairs[2].v)
		fmt.Println("[Three pairs have been stored in the sync.Map instance]")

		sMap.Range(func(key, value interface{}) bool {
			fmt.Printf("The result of an iteration in Range: %v, %v\n",
				key, value)
			return true
		})

		k0 := pairs[0].k
		v0, ok := sMap.Load(k0)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v0, ok, k0)

		k3 := pairs[3].k
		v3, ok := sMap.Load(k3)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v3, ok, k3)

		k2, v2 := pairs[2].k, pairs[2].v
		actual2, loaded2 := sMap.LoadOrStore(k2, v2)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual2, loaded2, k2, v2)
		v3 = pairs[3].v
		actual3, loaded3 := sMap.LoadOrStore(k3, v3)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual3, loaded3, k3, v3)

		k1 := pairs[1].k
		sMap.Delete(k1)
		fmt.Printf("[The pair with the key of %v has been removed from the sync.Map instance]\n",
			k1)
		v1, ok := sMap.Load(k1)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v1, ok, k1)
		v1 = pairs[1].v
		actual1, loaded1 := sMap.LoadOrStore(k1, v1)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual1, loaded1, k1, v1)

		sMap.Range(func(key, value interface{}) bool {
			fmt.Printf("The result of an iteration in Range: %v, %v\n",
				key, value)
			return true
		})
	}

}
