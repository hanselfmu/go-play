package cache

// LFUCache represents a least-frequently-used cache
// LeetCode algorithm challenge #460
type LFUCache struct {
	cap      int
	store    map[int]cacheValue
	freqList *freqNode
}

type cacheValue struct {
	val         int
	freqLine    *freqNode
	freqLinePos int
}

type freqNode struct {
	freq int
	next *freqNode
	keys []int
}

// Constructor creates a LFU cache
func Constructor(capacity int) LFUCache {
	freqList := freqNode{0, nil, []int{}}
	return LFUCache{capacity, map[int]cacheValue{}, &freqList}
}

// Get returns a value from LFU Cache
func (cache *LFUCache) Get(key int) int {
	if res := cache.store[key]; res.freqLine != nil {
		return res.val
	}

	return -1
}

// Put upserts a value into LFU Cache
func (cache *LFUCache) Put(key int, value int) {
	// checks if this key already exists; if so, updates the value
	if cur := cache.store[key]; cur.freqLine != nil {
		cur.val = value
	} else {
		// not existing currently; evict the LFRU and create a new one
		if len(cache.store) == cache.cap {
			// freqList head represents 0-used values; the next freqNode
			// represents the next smallest frequency
			freqList := cache.freqList
			if len(freqList.keys) > 0 {
				evictee := freqList.keys[0]
				
				freqList.keys = freqList.keys[1:len(freqList.keys)]
			} else {

			}
		}

		freqList := cache.freqList
		freqList.keys = append(freqList.keys, key)
		cache.store[key] = cacheValue{value, freqList, len(freqList.keys) - 1}
	}
}

/**
* Your LFUCache object will be instantiated and called as such:
* obj := Constructor(capacity);
* param_1 := obj.Get(key);
* obj.Put(key,value);

0 -> 2 -> 3 -> 5

For a Get to increase a cacheValue's frequency, we need to find the freqNode holding
this cacheValue, and and move this cacheValue to its next freqNode if the next
freqNode has the freq exactly 1 larger. Else a freqNode can be created.
*/
