package leetcode_75

// Given an integer array nums, find the subarray with the largest sum, and return its sum.
// A dynamic programming classic.
// We can do this with an O(n^2) algorithm filling in A[i, j] = sum of array starting at i, ending in j.
// Then search through the array to find the max value.
// I feel like this is actually a linear problem.
// Yes, memory serves me well.
/*
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
*/
// We can do something similar to the others, max sum before i, max sum after i.
// max sum after i is not required, we only need before.
// [-2,1,-3, 4,-1,2,1,-5,4]
// [0, 0, 1, 0, 4,3,5,6, 1]  (max of previous NOT including i)
// [-2, 1, 0, 4, 3, 5, 6,1,5] (max of previous INCLUDING i; solves n = 1 case)

func maxSubArray(nums []int) int {
	var maxSumBefore []int = make([]int, len(nums))
	var maxSoFar int
	for i, v := range nums {
		var j int
		if i == 0 {
			j = v
		} else {
			j = maxSumBefore[i - 1] + v
		}
		if j > 0 {
			maxSumBefore[i] = j
		}
		if j > maxSoFar {
			maxSoFar = j
		}
	}
	return maxSoFar
}
