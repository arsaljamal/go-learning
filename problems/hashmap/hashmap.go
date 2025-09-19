package main

import (
	"fmt"
)

// Entry represents a key-value pair
type Entry struct {
	key   string
	value string
	next  *Entry // linked list for collisions
}

// HashMap structure
type HashMap struct {
	buckets []*Entry
	size    int
}

// Create a new HashMap with given capacity
func NewHashMap(capacity int) *HashMap {
	return &HashMap{
		buckets: make([]*Entry, capacity),
		size:    capacity,
	}
}

// Simple hash function: sum of characters % size
func (h *HashMap) hash(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % h.size
}

// Put a key-value pair
func (h *HashMap) Put(key, value string) {
	index := h.hash(key)
	entry := h.buckets[index]

	// If bucket empty, insert directly
	if entry == nil {
		h.buckets[index] = &Entry{key: key, value: value}
		return
	}

	// If bucket has entries, check if key exists
	for entry != nil {
		if entry.key == key {
			entry.value = value // update
			return
		}
		if entry.next == nil {
			break
		}
		entry = entry.next
	}

	// Add new entry at the end of chain
	entry.next = &Entry{key: key, value: value}
}

// Get a value by key
func (h *HashMap) Get(key string) (string, bool) {
	index := h.hash(key)
	entry := h.buckets[index]

	for entry != nil {
		if entry.key == key {
			return entry.value, true
		}
		entry = entry.next
	}

	return "", false
}

func main() {
	hm := NewHashMap(5)

	hm.Put("HarryPotter", "Fantasy Book")
	hm.Put("SherlockHolmes", "Detective Book")
	hm.Put("Avengers", "Comic Book")

	// Retrieving
	if val, ok := hm.Get("HarryPotter"); ok {
		fmt.Println("HarryPotter:", val)
	}

	if val, ok := hm.Get("SherlockHolmes"); ok {
		fmt.Println("SherlockHolmes:", val)
	}

	if val, ok := hm.Get("Avengers"); ok {
		fmt.Println("Avengers:", val)
	}

	/*
	What this does:

	- Creates a HashMap with a fixed number of buckets.

	- Uses a hash function (sum of characters % size) to map keys to buckets.

	- If two keys land in the same bucket, it chains them in a linked list.

	 - Supports Put (insert/update) and Get (retrieve).
	*/
}

//HashMap = cupboard with drawers.

//Key = book name.

//Hash function = translator that turns names into drawer numbers.

//Value = book inside drawer.

//Collision = two books in the same drawer (solved with lists or shifting).

//Resize = buy a bigger cupboard when things get too full.


/* 
The Collision Problem

But what if two names give the same drawer number?

"HarryPotter" → 42

"SherlockHolmes" → 42

😱 Both want drawer 42!

Solutions:

1. Chaining (most common): Drawer #42 keeps a list of books.
 Both "HarryPotter" and "SherlockHolmes" live there. 
 When searching, we check the list.

2. Open addressing: If drawer 42 is full, go to 43, then 44,
 until you find an empty spot.

*/

/*
Resizing (Rehashing)

- If your cupboard (say 100 drawers) starts filling up,
 finding stuff gets slower because lists get long.

- So HashMap doubles the cupboard size (100 → 200 drawers).

- It re-distributes (rehashes) all items into new drawers.

- This keeps things fast.

*/