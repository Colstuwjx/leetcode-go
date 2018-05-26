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
	if !exist {
		return -1
	}

	return this.doublyLinkedList.values[idx].value
}

func (this *LRUCache) Iteral() map[int]int {
	values := make(map[int]int)
	for key, valueIdx := range this.valueMap {
		values[key] = this.doublyLinkedList.values[valueIdx].value
	}
	return values
}

func (this *LRUCache) evictKeyWithAppendTail(index, key, value int) *LRUCache {
	var appendIndex int

	// -1 means capacity is NOT full yet
	if index == -1 {
		appendIndex = len(this.valueMap)
	} else {
		// evict indexed item
		evictItem := this.doublyLinkedList.values[index]
		if evictItem.prior != nil {
			evictItem.prior.next = evictItem.next
		}

		if evictItem.next != nil {
			evictItem.next.prior = evictItem.prior
		}

		// head has been replaced, reset head
		if evictItem.index == this.doublyLinkedList.head.index {
			this.doublyLinkedList.head = this.doublyLinkedList.head.next
		}

		// reuse the evict item index
		appendIndex = index

		// clear evict item's key-value mapping
		// FIXME: delete operation is heavy, should be marked as unused
		delete(this.valueMap, evictItem.key)
	}

	newItem := &LinkedListItem{
		index: appendIndex,
		value: value,
		key:   key,
		prior: this.doublyLinkedList.tail,
		next:  nil,
	}

	if index == -1 {
		this.doublyLinkedList.values = append(this.doublyLinkedList.values, newItem)

		// init head if it is nil
		if this.doublyLinkedList.head == nil {
			this.doublyLinkedList.head = newItem
		}
	} else {
		this.doublyLinkedList.values[appendIndex] = newItem
	}

	// rotate tail if it exists
	if this.doublyLinkedList.tail != nil {
		this.doublyLinkedList.tail.next = newItem
	}

	this.doublyLinkedList.tail = newItem
	this.valueMap[key] = appendIndex
	return this
}

func (this *LRUCache) Put(key int, value int) {
	idx, exist := this.valueMap[key]
	if !exist {
		if len(this.valueMap) == this.capacity {
			evictHeadIndex := this.doublyLinkedList.head.index
			this.evictKeyWithAppendTail(evictHeadIndex, key, value)
		} else {
			this.evictKeyWithAppendTail(-1, key, value)
		}
	} else {
		this.evictKeyWithAppendTail(idx, key, value)
	}
}
