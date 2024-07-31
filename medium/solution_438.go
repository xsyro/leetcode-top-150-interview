package medium

import (
	"leetcode-top-150-interview-golang/easy"
)

type probability struct {
	combination string
	index       int
}

func FindAnagrams(s string, p string) []int {
	var anagramsIndexes []int
	var probabilities = []probability{}
	for i := 0; i < len(s)-(len(p)-1); i++ {
		probabilities = append(probabilities, probability{
			combination: s[i : i+len(p)],
			index:       i,
		})
	}
	// Using the method IsAnagram in Solution-242.go as a convenience
	for _, item := range probabilities {
		if easy.IsAnagram(item.combination, p) {
			anagramsIndexes = append(anagramsIndexes, item.index)
		}
	}
	return anagramsIndexes
}
