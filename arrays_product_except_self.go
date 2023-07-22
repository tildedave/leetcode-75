package leetcode_75

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/

// Obviously using division makes it trivial
// product[0] = nums[1] * nums[2] * nums[3]
// product[1] = nums[0] * nums[2] * nums[3]
// product[2] = nums[0] * nums[1] * nums[3]
// product[3] = nums[0] * nums[1] * nums[2]

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// OK so the idea is getting two arrays, product prefix / product suffix

// Prefix: [1, 1, 2, 6]
// Suffix: [24, 12, 4, 1]
// Then the result is prefix * suffix.

func productExceptSelf(nums []int) []int {
	var result []int = make([]int, len(nums))
	var n = len(nums) - 1

	var prefixSoFar int = 1
	for i, _ := range nums {
		if i == 0 {
			result[i] = 1
			continue
		}
		prefixSoFar *= nums[i - 1]
		result[i] = prefixSoFar
	}

	var suffixSoFar int = 1
	for i, _ := range nums {
		if i == 0 {
			result[n - i] *= 1
			continue
		}
		suffixSoFar *= nums[n - i + 1]
		result[n - i] *= suffixSoFar
	}

	return result
}
