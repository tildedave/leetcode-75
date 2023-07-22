package leetcode_75

/*
Given an integer array nums, find a
subarray
 that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.
*/

// So this is the same as the previous one, O(n) solution by recording max subarray up to this point and including the current item.
//
/*
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
*/

// [2, 6, -2, 4]

func maxProduct(nums []int) int {
	var maxProductSoFar []int = make([]int, len(nums))
	var maxSoFar int
	for i, v := range nums {
		var j int
		if i == 0 {
			j = v
		} else {
			j = v * maxProductSoFar[i - 1]
		}
		if j > v {
			maxProductSoFar[i] = j
		} else {
			maxProductSoFar[i] = v
		}
		if maxProductSoFar[i] > maxSoFar {
			maxSoFar = maxProductSoFar[i]
		}
	}
	return maxSoFar
}
