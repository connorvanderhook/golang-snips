package structures

import (
	"fmt"
	"sync"
)

type HashTable struct {
	items map[int]string
	lock  sync.RWMutex
}

// the hash() private function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
func hash(k string) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

// Put in da htable
func (ht *HashTable) Put(k string, v string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := hash(k)
	if ht.items == nil {
		ht.items = make(map[int]string)
	}
	ht.items[i] = v
}

// Delete from the table
func (ht *HashTable) Delete(k string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := hash(k)
	delete(ht.items, i)
}
