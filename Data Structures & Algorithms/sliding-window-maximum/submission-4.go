type Item struct {
	value int
	index int
}

type MaxHeap []Item

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}


func maxSlidingWindow(nums []int, k int) []int {
    var ans []int
	n := len(nums)
	l := 0
	h := &MaxHeap{}
	heap.Init(h)
	for r := 0; r < k; r++ {
		heap.Push(h, Item{value: nums[r], index: r})
	}
	mx := (*h)[0]
	ans = append(ans, mx.value)

	for r := k; r < n; r++ {
		l++
		heap.Push(h, Item{value: nums[r], index: r})
		mx = (*h)[0]
		for mx.index < l {
			heap.Pop(h)
			mx = (*h)[0]
		}
		
		ans = append(ans, mx.value)
	}
	return ans
}
