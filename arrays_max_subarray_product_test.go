package leetcode_75

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMaxSubArrayProduct(t *testing.T) {
	assert.Equal(t, 6, maxProduct([]int{2,3,-2,4}))
	assert.Equal(t, 0, maxProduct([]int{-2,0,-1}))
	assert.Equal(t, 12, maxProduct([]int{12}))
}
