package main

/*
   Description

   Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

   Note:

   The number of elements initialized in nums1 and nums2 are m and n respectively.
   You may assume that nums1 has enough space (size that is greater or equal to m + n)
   to hold additional elements from nums2.

   Example:

   ```
   Input:
   nums1 = [1,2,3,0,0,0], m = 3
   nums2 = [2,5,6],       n = 3

   Output: [1,2,2,3,5,6]
   ```
*/

/*
   解题思路:

   和[problem 21](https://leetcode.com/problems/merge-two-sorted-lists/description/)有所区别的是
   这里的两个list是用数组来存储的，因此剪切的操作可能不再适用，而且需要注意的是最终结果是得存储在nums1的

   1, 2, 3, 5, 7
   2, 5, 6

   不妨考虑倒序插入，数组的两个索引m、n，每次判断最大值哪个大，大的那个，插入到最后面，然后该序列索引值-1

   0. m = 5, n = 3
   1. 7 > 6, a[m+n-1] = 7, m--, m = 4
   2. 5 < 6, a[m+n-1] = 6, n--, n = 2
   3. ...
*/

import "log"

func merge(nums1 []int, m int, nums2 []int, n int) {
	midx := m
	nidx := n

	for {
		if midx <= 0 && nidx <= 0 {
			break
		}

		if midx == 0 {
			nums1[midx+nidx-1] = nums2[nidx-1]
			nidx = nidx - 1
			continue
		}

		if nidx == 0 {
			nums1[midx+nidx-1] = nums1[midx-1]
			midx = midx - 1
			continue
		}

		if nums1[midx-1] > nums2[nidx-1] {
			nums1[midx+nidx-1] = nums1[midx-1]
			midx = midx - 1
		} else {
			nums1[midx+nidx-1] = nums2[nidx-1]
			nidx = nidx - 1
		}
	}
}

func main() {
	var (
		nums1 = []int{1, 2, 3, 5, 7, 0, 0, 0}
		nums2 = []int{2, 5, 6}
	)

	merge(nums1, 5, nums2, 3)
	log.Println("Merge result: ", nums1)
}
