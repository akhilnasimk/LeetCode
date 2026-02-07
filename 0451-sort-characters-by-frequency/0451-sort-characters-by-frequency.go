type Item struct {
    ch    rune
    count int
}

type baseheap []Item

func (h baseheap) Len() int { return len(h) }
func (h baseheap) Less(i, j int) bool {
    return h[i].count > h[j].count // max heap
}
func (h baseheap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *baseheap) Push(x interface{}) {
    *h = append(*h, x.(Item))
}

func (h *baseheap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func frequencySort(s string) string {
    // 1. Count frequency
    freq := make(map[rune]int)
    for _, ch := range s {
        freq[ch]++
    }

    // 2. Build heap
    h := make(baseheap, 0, len(freq))
    heap.Init(&h)

    for ch, cnt := range freq {
        heap.Push(&h, Item{ch: ch, count: cnt})
    }

    // 3. Build result efficiently
    var builder strings.Builder
    builder.Grow(len(s)) // optional but nice

    for h.Len() > 0 {
        item := heap.Pop(&h).(Item)
        for i := 0; i < item.count; i++ {
            builder.WriteRune(item.ch)
        }
    }

    return builder.String()
}
