package top_k_frequent_words

import (
	"container/heap"
	"strings"
)

type WordFrequency struct {
	Word  string
	Count int
}

type MinHeap []WordFrequency

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].Count == h[j].Count {
		return strings.Compare(h[i].Word, h[j].Word) > 0
	}
	return h[i].Count < h[j].Count
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(WordFrequency))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	wordFrequencies := make(map[string]int)
	for _, word := range words {
		wordFrequencies[word]++
	}

	h := &MinHeap{}
	heap.Init(h)
	for word, count := range wordFrequencies {
		heap.Push(h, WordFrequency{word, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[k-i-1] = heap.Pop(h).(WordFrequency).Word
	}

	return result
}
