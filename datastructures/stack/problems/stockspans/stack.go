package stockspans

import "bubl/study/datastructures/stack"

// SpansByStack get sppans of a stock by using stack
func SpansByStack(stocks []int) []int {
	st := stack.NewStack(len(stocks))
	heights := make([]int, len(stocks))
	for i := range stocks {
		for !st.IsEmpty() {
			idx, _ := st.Top()
			if stocks[i] >= stocks[idx] {
				st.Pop()
			} else {
				st.Push(i)
				heights[i] = idx
				break
			}
		}
		if st.IsEmpty() {
			heights[i] = -1
			st.Push(i)
			continue
		}
	}

	spans := make([]int, len(heights))
	for i, h := range heights {
		spans[i] = i - h
	}

	return spans
}
