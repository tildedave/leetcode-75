package leetcode_75

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.

log n says we need to do a tree-type approach, descending through the array and eliminating half of it roughly at a time.
How can we "rule out" portions of the array?
We know the array is guaranteed to be sorted.

*/

func findMin(nums []int) int {
	minIndex := 0
	maxIndex := len(nums) - 1
	for maxIndex > minIndex {
		var i int = (maxIndex + minIndex) / 2
		// Base cases; are we at the end.
		if nums[minIndex] < nums[maxIndex] {
			return nums[minIndex]
		} else if i + 1 < len(nums) - 1 && nums[i] > nums[i + 1] {
			return nums[i + 1]
		}
		// Otherwise our pivot is less than its next neighbor so we need to winnow the array.
		if nums[i] > nums[maxIndex] {
			minIndex = i
		} else if nums[minIndex] > nums[i] {
			maxIndex = i
		}
	}
	return nums[maxIndex]
}
