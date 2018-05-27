package main

/*
   <Level> easy

   Description

   Merge two sorted linked lists and return it as a new list.
   The new list should be made by splicing together the nodes of the first two lists.

   Example:

   ```
   Input: 1->2->4, 1->3->4
   Output: 1->1->2->3->4->4
   ```
*/

/*
   解题思路:

   因为实现方式是需要两个链表接片起来
   设两个链表为A和B，那么每次取出链表里的两个元素比较
   每次判断大小，如果某个元素小的话，走到最后一个比前面元素小的指针，然后将前面的子串剪切到前面元素上
   （因为比前面元素小，而且又是排序好的，因此子串肯定是全局最小的）
   循环往复，直到有一个串走到尽头，即判断结束

   1. 1 <= 1 and 2 > 1, l1: 2->4, l2: 1->1->3->4, m = 2, n = 1
   2. 1 < 2 and 3 > 2, l1: 1->1->2->4, l2: 3->4, m = 2, n = 3
   3. 2 < 3 and 4 > 3, l1: 4, l2: 1->1->2->3->4, m = 4, n = 3
   4. 4 > 3 and 4 >= 4, l1: 1->1->2->3->4->4, l2: nil
   5. 结束
*/

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	m := l1
	n := l2

	// set the head
	if m.Val <= n.Val {
		head = m
	} else {
		head = n
	}

	for {
		if m == nil || n == nil {
			break
		}

		if m.Val <= n.Val {
			for {
				if m.Next != nil && m.Next.Val <= n.Val {
					m = m.Next
				} else {
					m.Next, m = n, m.Next
					break
				}
			}
		} else {
			for {
				if n.Next != nil && n.Next.Val <= m.Val {
					n = n.Next
				} else {
					n.Next, n = m, n.Next
					break
				}
			}
		}
	}

	return head
}

func main() {
	var l1, l2 *ListNode
	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 = &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	origin := mergeTwoLists(l1, l2)
	for {
		if origin != nil {
			log.Print(origin.Val, " -> ")
			origin = origin.Next
		} else {
			break
		}
	}
}
