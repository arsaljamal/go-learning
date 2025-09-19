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
	count   int // number of elements
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

// Load factor = count / size
func (h *HashMap) loadFactor() float64 {
	return float64(h.count) / float64(h.size)
}

// Put a key-value pair
func (h *HashMap) Put(key, value string) {
	index := h.hash(key)
	entry := h.buckets[index]

	// If bucket empty, insert directly
	if entry == nil {
		h.buckets[index] = &Entry{key: key, value: value}
		h.count++
	} else {
		// Check if key exists
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
		h.count++
	}

	// Rehash if load factor > 0.7
	if h.loadFactor() > 0.7 {
		h.rehash()
	}
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

// Rehash (double the size and redistribute)
func (h *HashMap) rehash() {
	newSize := h.size * 2
	newBuckets := make([]*Entry, newSize)

	oldBuckets := h.buckets
	h.buckets = newBuckets
	h.size = newSize
	h.count = 0 // will be recalculated as we insert again

	for _, bucket := range oldBuckets {
		entry := bucket
		for entry != nil {
			h.Put(entry.key, entry.value) // reinsert into new table
			entry = entry.next
		}
	}

	fmt.Println("Rehashed! New size:", newSize)
}

func main() {
	hm := NewHashMap(4)

	hm.Put("HarryPotter", "Fantasy Book")
	hm.Put("SherlockHolmes", "Detective Book")
	hm.Put("Avengers", "Comic Book")
	hm.Put("Inception", "Sci-Fi Movie") // this should trigger rehash

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
	if val, ok := hm.Get("Inception"); ok {
		fmt.Println("Inception:", val)
	}
}
