package leetcode_75

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

*/

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4

// Seems like we can find the minimum and then based on that search both sides
// of the tree which are then sorted.
// Is there a faster way?
// It's not clear that we can know which side of the tree to discard by looking
// at a pivot.  So finding the break point seems required.


func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	pivotIdx := len(nums) / 2
	r := nums[pivotIdx]
	if r == target {
		return pivotIdx
	}
	if r < target {
		res := binarySearch(nums[pivotIdx + 1:], target)
		if res == -1 {
			return -1
		}
		return pivotIdx + 1 + res
	}
	return binarySearch(nums[0:pivotIdx], target)
}

func search(nums []int, target int) int {
	minIndex := 0
	maxIndex := len(nums) - 1
	var splitIndex int
	for maxIndex > minIndex {
		if nums[minIndex] < nums[maxIndex] {
			splitIndex = minIndex
			break
		}
		n := (minIndex + maxIndex) / 2
		if nums[n] > nums[maxIndex] {
			minIndex = n + 1
		} else {
			maxIndex = n
		}
		// This is susceptible to an infinite loop in the case of malformed
		// input, we should be smarter.
	}
	lower := nums[0:splitIndex]
	upper := nums[splitIndex:]

	idx := binarySearch(lower, target)
	if idx >= 0 {
		return idx
	}
	idx = binarySearch(upper, target)
	if idx >= 0 {
		return idx + splitIndex
	}
	return -1
}
