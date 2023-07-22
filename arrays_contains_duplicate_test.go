package leetcode_75

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	assert.True(t, containsDuplicate([]int{1,2,3,1}))
	assert.False(t, containsDuplicate([]int{1,2,3,4}))
	assert.True(t, containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2}))
}
