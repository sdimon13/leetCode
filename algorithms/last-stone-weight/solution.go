package last_stone_weight

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		y := heap.Pop(&h).(int)
		x := heap.Pop(&h).(int)

		if x != y {
			heap.Push(&h, y-x)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return heap.Pop(&h).(int)
}
