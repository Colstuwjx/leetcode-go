package main

/*
   Description

   Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

   get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
   put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

   Follow up:
   Could you do both operations in O(1) time complexity?

   Example:

   ```
   LRUCache cache = new LRUCache( 2 // capacity );

   cache.put(1, 1);
   cache.put(2, 2);
   cache.get(1);       // returns 1
   cache.put(3, 3);    // evicts key 2
   cache.get(2);       // returns -1 (not found)
   cache.put(4, 4);    // evicts key 1
   cache.get(1);       // returns -1 (not found)
   cache.get(3);       // returns 3
   cache.get(4);       // returns 4
   ```
*/

/*
   解题思路：

   因为有O(1)的算法复杂度要求，所以需要用到hashmap
   每次添加数据时，用一个双向链表来存储对应的value，同时用链表建立前后更新关系，用hashmap存储key和对应链表数组item的索引；
   每次获取数据时，根据hashmap获取到的item索引取出链表数组里的实际value；
   每次容量填满需要踢出数据时（即hashmap key not exist时），去掉最久没更新的一个item数据（即链表的头），并在尾部插入一条新数据；
*/

import "log"

const (
	NotExistIndex = -1
)

type LinkedListItem struct {
	index int
	key   int
	value int
	prior *LinkedListItem
	next  *LinkedListItem
}

type DoublyLinkedList struct {
	head   *LinkedListItem
	tail   *LinkedListItem
	values []*LinkedListItem
}

type LRUCache struct {
	doublyLinkedList *DoublyLinkedList
	valueMap         map[int]int
	capacity         int
}

func Constructor(capacity int) LRUCache {
	var valueSlices []*LinkedListItem

	if capacity <= 0 {
		panic("capacity must be positive number!")
	}

	return LRUCache{
		doublyLinkedList: &DoublyLinkedList{
			head:   nil,
			tail:   nil,
			values: valueSlices,
		},
		valueMap: map[int]int{},
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	idx, exist := this.valueMap[key]
	if !exist || idx == -1 {
		return -1
	}

	value := this.doublyLinkedList.values[idx].value
	this.evictKeyWithAppendTail(idx, key, value)
	return value
}

func (this *LRUCache) IndexMap() map[int]int {
	return this.valueMap
}

func (this *LRUCache) Iteral() [][]int {
	var values [][]int
	pointer := this.doublyLinkedList.head
	for {
		if pointer == nil {
			break
		}

		values = append(values, []int{pointer.key, pointer.value})
		pointer = pointer.next
	}

	return values
}

func (this *LRUCache) evictKeyWithAppendTail(index, key, value int) *LRUCache {
	item := &LinkedListItem{
		index: index,
		key:   key,
		value: value,
		prior: nil,
		next:  nil,
	}

	// 1. head is nil or cap = 1
	if this.doublyLinkedList.head == nil || this.capacity == 1 {
		currentHead := this.doublyLinkedList.head
		if currentHead != nil && currentHead.key != item.key {
			this.valueMap[currentHead.key] = NotExistIndex
		}

		item.index = 0
		this.doublyLinkedList.head = item
		this.doublyLinkedList.tail = nil
		if len(this.doublyLinkedList.values) != 0 {
			this.doublyLinkedList.values[0] = item
		} else {
			this.doublyLinkedList.values = append(this.doublyLinkedList.values, item)
		}

		this.valueMap[key] = item.index
		return this
	}

	// 2. head is not nil, and need to evict head
	if this.doublyLinkedList.head.index == index {
		// ONLY one item case
		if this.doublyLinkedList.head.next == nil {
			this.doublyLinkedList.head = item
			this.doublyLinkedList.tail = nil

			evictItemKey := this.doublyLinkedList.values[index].key
			this.valueMap[evictItemKey] = NotExistIndex
			this.doublyLinkedList.values[index] = item
			this.valueMap[key] = index
			return this
		}

		currentHead := this.doublyLinkedList.head
		this.doublyLinkedList.head.next.prior = nil
		this.doublyLinkedList.head = currentHead.next
		this.doublyLinkedList.tail.next = item
		item.prior = this.doublyLinkedList.tail
		this.doublyLinkedList.tail = item

		evictHeadKey := this.doublyLinkedList.values[index].key
		this.valueMap[evictHeadKey] = NotExistIndex
		this.doublyLinkedList.values[index] = item
		this.valueMap[key] = item.index
		return this
	}

	// 3. tail is nil and need to init
	if this.doublyLinkedList.tail == nil {
		this.doublyLinkedList.head.next = item
		this.doublyLinkedList.tail = item
		this.doublyLinkedList.tail.prior = this.doublyLinkedList.head
		this.doublyLinkedList.values = append(this.doublyLinkedList.values, item)
		this.valueMap[key] = index
		return this
	}

	// 4. evict other node(include tail), or append item
	if len(this.doublyLinkedList.values) > index {
		evictItem := this.doublyLinkedList.values[index]
		if evictItem.prior.next != nil {
			evictItem.prior.next = evictItem.next
		} else {
			this.doublyLinkedList.head.next = item
		}

		if evictItem.next != nil {
			evictItem.next.prior = evictItem.prior
		} else {
			this.doublyLinkedList.tail = evictItem.prior
		}
		this.valueMap[evictItem.key] = NotExistIndex
	}

	item.prior = this.doublyLinkedList.tail
	this.doublyLinkedList.tail.next = item
	this.doublyLinkedList.tail = item
	if len(this.doublyLinkedList.values) == index {
		this.doublyLinkedList.values = append(this.doublyLinkedList.values, item)
	} else {
		this.doublyLinkedList.values[index] = item
	}

	this.valueMap[key] = item.index
	return this
}

func (this *LRUCache) Put(key int, value int) {
	idx, exist := this.valueMap[key]
	if !exist || idx == NotExistIndex {
		if len(this.doublyLinkedList.values) >= this.capacity {
			evictHeadIndex := this.doublyLinkedList.head.index
			this.evictKeyWithAppendTail(evictHeadIndex, key, value)
		} else {
			appendToTailIndex := len(this.doublyLinkedList.values)
			this.evictKeyWithAppendTail(appendToTailIndex, key, value)
		}
	} else {
		this.evictKeyWithAppendTail(idx, key, value)
	}
}

// func main() {
// 	cache := Constructor(1)
// 	log.Println("Put (2, 1)")
// 	cache.Put(2, 1)

// 	log.Println("Geting 2")
// 	log.Println("Get 2: ", cache.Get(2))

// 	log.Println("Put (3, 2)")
// 	cache.Put(3, 2)

// 	log.Println("Geting 2")
// 	log.Println("Get 2: ", cache.Get(2))

// 	log.Println("Geting 3")
// 	log.Println("Get 3: ", cache.Get(3))

// 	log.Println("Test passed!")
// }
