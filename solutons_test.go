package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
func TestSolution27(t *testing.T) {
	assert.Equal(t, 2, removeElement([]int{3, 2, 2, 3}, 3))
	assert.Equal(t, 5, removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
