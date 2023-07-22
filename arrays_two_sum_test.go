package leetcode_75

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestTwoSum_1(t *testing.T) {
	input := []int{2,7,11,15}
	assert.Equal(t, []int{0, 1}, twoSum(input, 9))
}

func TestTwoSum_2(t *testing.T) {
	input := []int{3,2,4}
	assert.Equal(t, []int{1, 2}, twoSum(input, 6))
}

func TestTwoSum_3(t *testing.T) {
	input := []int{3, 3}
	assert.Equal(t, []int{0, 1}, twoSum(input, 6))
}
