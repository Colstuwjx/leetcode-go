package main

/*
   <Level> easy

   Description

   Given a sorted array and a target value, return the index if the target is found.
   If not, return the index where it would be if it were inserted in order.
   You may assume no duplicates in the array.

   Example:

   ```
   Example 1:

   Input: [1,3,5,6], 5
   Output: 2
   Example 2:

   Input: [1,3,5,6], 2
   Output: 1
   Example 3:

   Input: [1,3,5,6], 7
   Output: 4
   Example 4:

   Input: [1,3,5,6], 0
   Output: 0
   ```
*/

/*
   解题思路:

   这是一道二分查找的变种题，在二分搜索的同时，如果超出范围，或者不匹配值时做相应处理即可.
   在high和low的范围缩小到差距为1时，进行判断是值匹配还是需要做插入操作时对应的索引，插入优先放在后面也就是索引+1
*/

import "log"

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	if target > nums[high] {
		return high + 1
	}

	if target < nums[low] {
		return 0
	}

	for {
		if high-low <= 1 {
			if target > nums[high] {
				return high + 1
			}

			if target == nums[high] {
				return high
			}

			if target < nums[high] && target > nums[low] {
				return low + 1
			}

			if target <= nums[low] {
				return low
			}
		}

		mid := (high + low) / 2
		if target == nums[mid] {
			return mid
		}

		if target > nums[mid] {
			low = mid + 1
			continue
		}

		if target < nums[mid] {
			high = mid - 1
			continue
		}
	}

	return -1
}

func main() {
	log.Println("target result: ", searchInsert([]int{1, 3, 5, 6}, 5))
}
