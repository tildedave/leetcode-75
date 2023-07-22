package leetcode_75

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 1, findMin([]int{3,4,5,1,2}))
	assert.Equal(t, 0, findMin([]int{4,5,6,7,0,1,2}))
	assert.Equal(t, 11, findMin([]int{11,13,15,17}))
}
