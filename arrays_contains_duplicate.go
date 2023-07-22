package leetcode_75

func containsDuplicate(nums []int) bool {
	seenMap := make(map[int]int)
	for _, v := range nums {
		if seenMap[v] != 0 {
			return true
		} else {
			seenMap[v] = 1
		}
	}
	return false
}
