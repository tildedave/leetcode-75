package leetcode_75

/**
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Input: nums = [3,2,4], target = 6
Output: [1,2]

Input: nums = [3,3], target = 6
Output: [0,1]
*/

func twoSum(nums []int, target int) []int {
	adjusted := make(map[int]int)
	for i, v := range nums {
		adjusted[target - v] = i
	}
	var results []int
	for i, v := range nums {
		if adjusted[v] > 0 {
			return []int{i, adjusted[v]}
		}
	}
	return results
}
