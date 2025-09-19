package main

import "fmt"

type Node struct {
	key,value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache map[int]*Node
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	// Case 1: Key exists, move it to the front
	//  and return its value.
	if node,ok := this.cache[key]; ok {
		this.moveToFront(node)
		return  node.value
	}
	// Case 2: Key does not exist, return -1.
	return  -1
}

func (this *LRUCache) Put(key int, value int) {
	//Case 1: Key already exists, 
	// update the value and move it to the front.
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToFront(node)
		return
	}
	//Case 2: capacity is reached,
	// remove the least recently used item
	//  (tail's previous node).
	// remove it from the cache map as well.
	if len(this.cache) == this.capacity {
		lru := this.tail.prev
		this.remove((lru))
		delete(this.cache, lru.key)
	}
	
	//Case 3: Key does not exist, create a new node,
	// insert it at the front, and add it to the cache map.
	newNode := &Node{
		key: key, 
		value: value,
	}
	this.insertAfterHead(newNode)
	this.cache[key] = newNode
}

func (this *LRUCache) moveToFront(node *Node) {
	// Remove node from its current position.
	// Reinsert it right after the head.
	this.remove(node)
	this.insertAfterHead(node)
}

func (this *LRUCache) remove(node *Node) {
	// Bypass the node to remove it from the list.
	// Update the previous and next pointers.
	node.prev.next = node.prev
	node.next.prev = node.next
}

func (this *LRUCache) insertAfterHead(node *Node) {
	// Insert the node right after the head.
	// Update the previous and next pointers accordingly.
	// Adjust the head's next node's previous pointer.
	// Adjust the head's next pointer to point
	// to the new node.
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}
 
func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // returns 1
	lru.Put(3, 3) // removes key 2
	fmt.Println(lru.Get(2)) // returns -1 (not found)
}