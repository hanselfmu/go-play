package cache

import (
	"fmt"
	"testing"
)

func TestLFUCache(t *testing.T) {
	instructions := []string{"put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	puts := [][2]int{{10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {2, 19}, {5, 25}, {9, 22}, {5, 5}, {1, 30}, {9, 12}, {4, 30}, {9, 3}, {6, 14}, {3, 1}, {10, 11}, {2, 14}, {11, 4}, {12, 24}, {5, 18}, {7, 23}, {3, 27}, {2, 12}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {9, 29}, {8, 21}, {6, 30}, {1, 12}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {3, 4}, {11, 30}, {4, 29}, {3, 4}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {13, 17}, {2, 27}, {11, 15}, {9, 19}, {2, 15}, {3, 16}, {12, 17}, {9, 1}, {6, 19}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}}
	gets := [][1]int{{13}, {2}, {3}, {8}, {11}, {7}, {5}, {8}, {9}, {9}, {10}, {10}, {3}, {8}, {1}, {5}, {4}, {13}, {8}, {12}, {5}, {6}, {5}, {10}, {5}, {12}, {3}, {9}, {6}, {1}, {10}, {3}, {8}, {7}, {5}, {12}, {1}, {4}, {5}, {5}, {1}}

	idxP, idxG := 0, 0
	cache := Constructor(10)
	for _, ins := range instructions {
		if ins == "put" {
			fmt.Printf("PUT %v\n", puts[idxP])
			cache.Put(puts[idxP][0], puts[idxP][1])
			idxP++
		} else {
			fmt.Printf("GET %v: ", gets[idxG])
			fmt.Printf(" %d\n", cache.Get(gets[idxG][0]))
			idxG++
		}

		freqLine := cache.freqList
		for freqLine != nil {
			fmt.Printf("freq %d, size %d: ", freqLine.freq, freqLine.keyListSize)

			val := freqLine.keyListHead
			for val != nil {
				fmt.Printf("%v\t", val.key)
				val = val.next
			}

			fmt.Printf("\n")
			freqLine = freqLine.next
		}
	}
}
