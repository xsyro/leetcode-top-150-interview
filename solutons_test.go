package leetcode_top_150_interview_golang

import (
	"github.com/stretchr/testify/assert"
	"leetcode-top-150-interview-golang/easy"
	"leetcode-top-150-interview-golang/medium"
	"testing"
)

// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
func TestSolution27(t *testing.T) {
	assert.Equal(t, 2, easy.RemoveElement([]int{3, 2, 2, 3}, 3))
	assert.Equal(t, 5, easy.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

// https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
func TestSolution88(t *testing.T) {
	t.Run("TestSolution88", func(t *testing.T) {
		easy.Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	})
}

func TestSolution151(t *testing.T) {
	assert.Equal(t, medium.ReverseWords("a good   example"), "example good a")
	assert.Equal(t, medium.ReverseWords("  hello world  "), "world hello")
	assert.Equal(t, medium.ReverseWords("the sky is blue"), "blue is sky the")
}

func TestSolution438(t *testing.T) {
	assert.Equal(t, medium.FindAnagrams("cbaebabacd", "abc"), []int{0, 6})
	assert.Equal(t, medium.FindAnagrams("abab", "ab"), []int{0, 1, 2})
}
