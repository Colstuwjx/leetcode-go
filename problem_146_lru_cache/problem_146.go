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
   解题思路:

   因为有O(1)的算法复杂度要求，所以需要用到hashmap
   每次添加数据时，用一个双向链表来存储对应的value，同时用链表建立前后更新关系，用hashmap存储key和对应链表数组item的索引；
   每次获取数据时，根据hashmap获取到的item索引取出链表数组里的实际value；
   每次容量填满需要踢出数据时（即hashmap key not exist时），去掉最久没更新的一个item数据（即链表的头），并在尾部插入一条新数据；

   改进:

   大体思路仍然不变，改进点在于实现LRUCache的removeOldest方法来专门管理旧节点的过期，允许头尾相连，减少边际条件的判断
   之前的解法：[Ugly but works solution](https://github.com/Colstuwjx/leetcode-go/blob/c977440992103edf75eabd05ac89f7e1d98eff2e/146_LRU_Cache/problem_146.go)
*/

import "log"

type LinkedListNode struct {
	key   int
	value int
	prev  *LinkedListNode
	next  *LinkedListNode
}

func NewNode(key, value int) *LinkedListNode {
	return &LinkedListNode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

type DoublyLinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
}

func (this *DoublyLinkedList) Insert(node *LinkedListNode) *DoublyLinkedList {
	if this.head == nil {
		this.head, this.tail = node, node
		node.prev, node.next = nil, nil
		return this
	}

	this.tail.next = node
	node.prev = this.tail
	this.tail = node
	return this
}

func (this *DoublyLinkedList) RemoveHead() *DoublyLinkedList {
	this.Remove(this.head)
	return this
}

func (this *DoublyLinkedList) Remove(node *LinkedListNode) *DoublyLinkedList {
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
		return this
	}

	if node == this.head {
		node.next.prev = nil
		this.head = node.next
		return this
	}

	if node == this.tail {
		node.prev.next = nil
		this.tail = node.prev
		return this
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	return this
}

type LRUCache struct {
	doublyLinkedList *DoublyLinkedList
	valueMap         map[int]*LinkedListNode
	capacity, length int
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		panic("capacity must be positive number!")
	}

	return LRUCache{
		doublyLinkedList: &DoublyLinkedList{
			head: nil,
			tail: nil,
		},
		valueMap: map[int]*LinkedListNode{},
		capacity: capacity,
		length:   0,
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.valueMap[key]
	if !exist {
		return -1
	}

	this.doublyLinkedList.Remove(node)
	this.doublyLinkedList.Insert(node)
	return node.value
}

func (this *LRUCache) IndexMap() map[int]int {
	values := make(map[int]int)
	for key, node := range this.valueMap {
		values[key] = node.value
	}
	return values
}

func (this *LRUCache) Iterate() [][]int {
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

func (this *LRUCache) Put(key int, value int) {
	_, exist := this.valueMap[key]
	if !exist {
		node := NewNode(key, value)
		this.valueMap[key] = node
		this.doublyLinkedList.Insert(node)
		this.length += 1
		if this.length > this.capacity {
			this.length -= 1
			delete(this.valueMap, this.doublyLinkedList.head.key)
			this.doublyLinkedList.RemoveHead()
		}
	} else {
		this.doublyLinkedList.Remove(this.valueMap[key])
		this.doublyLinkedList.Insert(this.valueMap[key])
		this.valueMap[key].value = value
	}
}

func main() {
	cache := Constructor(1)
	log.Println("Put (2, 1)")
	cache.Put(2, 1)

	log.Println("Geting 2")
	log.Println("Get 2: ", cache.Get(2))

	log.Println("Put (3, 2)")
	cache.Put(3, 2)

	log.Println("Geting 2")
	log.Println("Get 2: ", cache.Get(2))

	log.Println("Geting 3")
	log.Println("Get 3: ", cache.Get(3))

	log.Println("Test passed!")
}
