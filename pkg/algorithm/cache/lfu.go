package cache

// LFUCache represents a least-frequently-used cache
// LeetCode algorithm challenge #460
type LFUCache struct {
	cap      int
	store    map[int]cacheValue
	freqList *freqNode
}

type cacheValue struct {
	val           int
	freqLine      *freqNode
	keyInFreqLine *cacheKeyNode
}

type freqNode struct {
	freq        int
	next        *freqNode
	keyListSize int
	keyListHead *cacheKeyNode
	keyListTail *cacheKeyNode
}

func (node *freqNode) removeHead() {
	if node.keyListSize <= 1 {
		node.keyListHead = nil
		node.keyListTail = nil
		node.keyListSize = 0
	} else {
		oldHead := node.keyListHead
		node.keyListHead = oldHead.next
		oldHead.deleteSelf()
		node.keyListSize--
	}
}

func (node *freqNode) appendTail(newKey *cacheKeyNode) {
	if node.keyListSize > 0 {
		node.keyListTail.insertAfter(newKey)
	}

	if node.keyListHead == nil {
		node.keyListHead = newKey
	}
	node.keyListTail = newKey
	node.keyListSize++
}

type cacheKeyNode struct {
	key  int
	prev *cacheKeyNode
	next *cacheKeyNode
}

func (node *cacheKeyNode) insertBefore(newNode *cacheKeyNode) {
	if node.prev != nil {
		node.prev.next = newNode
		newNode.prev = node.prev
	} else {
		newNode.prev = nil
	}
	node.prev = newNode
	newNode.next = node
}

func (node *cacheKeyNode) insertAfter(newNode *cacheKeyNode) {
	if node.next != nil {
		node.next.prev = newNode
		newNode.next = node.next
	} else {
		newNode.next = nil
	}
	node.next = newNode
	newNode.prev = node
}

func (node *cacheKeyNode) deleteSelf() {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node = nil
}

// Constructor creates a LFU cache
func Constructor(capacity int) LFUCache {
	freqList := freqNode{0, nil, 0, nil, nil}
	return LFUCache{capacity, map[int]cacheValue{}, &freqList}
}

func (cache *LFUCache) addFreq(key int) {
	if res := cache.store[key]; res.freqLine != nil {
		// move this cacheValue from its current freqNode to the next freqNode;
		// if the next freqNode's frequency is exactly larger by 1, append to the next freqNode
		// else create a new freqNode and append it there
		curFreqNode := res.freqLine

		if curFreqNode.keyListHead == res.keyInFreqLine {
			curFreqNode.keyListHead = res.keyInFreqLine.next
		}
		if curFreqNode.keyListTail == res.keyInFreqLine {
			curFreqNode.keyListTail = res.keyInFreqLine.prev
		}
		res.keyInFreqLine.deleteSelf()
		curFreqNode.keyListSize--
		newKeyNode := &cacheKeyNode{key, nil, nil}

		if nextFreqNode := curFreqNode.next; nextFreqNode == nil {
			// no next frequency; create a new one and append this key to its keyList
			newFreqNode := &freqNode{curFreqNode.freq + 1, nil, 1, newKeyNode, newKeyNode}
			curFreqNode.next = newFreqNode
			res.freqLine = newFreqNode
		} else if nextFreqNode.freq == curFreqNode.freq+1 {
			res.freqLine = nextFreqNode
			nextFreqNode.appendTail(newKeyNode)
		} else {
			newFreqNode := &freqNode{curFreqNode.freq + 1, nextFreqNode, 1, newKeyNode, newKeyNode}
			curFreqNode.next = newFreqNode
			res.freqLine = newFreqNode
		}

		res.keyInFreqLine = newKeyNode
		cache.store[key] = res
	}
}

// Get returns a value from LFU Cache
func (cache *LFUCache) Get(key int) int {
	if cache.cap == 0 {
		return -1
	}

	if res := cache.store[key]; res.freqLine != nil {
		cache.addFreq(key)
		return res.val
	}

	return -1
}

// Put upserts a value into LFU Cache
func (cache *LFUCache) Put(key int, value int) {
	if cache.cap == 0 {
		return
	}

	// checks if this key already exists; if so, updates the value
	if cur := cache.store[key]; cur.freqLine != nil {
		cur.val = value
		cache.store[key] = cur
		cache.addFreq(key)
	} else {
		freqList := cache.freqList
		// not existing currently; evict the LFRU and create a new one
		if len(cache.store) == cache.cap {
			// freqList head represents 0-used values; the next freqNode
			// represents the next smallest frequency
			removed := false
			curFreqNode := freqList
			for !removed {
				if curFreqNode.keyListSize > 0 {
					delete(cache.store, curFreqNode.keyListHead.key)
					curFreqNode.removeHead()
					removed = true
				}
				curFreqNode = curFreqNode.next
			}
		}

		newKeyNode := &cacheKeyNode{key, freqList.keyListTail, nil}
		freqList.appendTail(newKeyNode)
		cache.store[key] = cacheValue{value, freqList, newKeyNode}
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
