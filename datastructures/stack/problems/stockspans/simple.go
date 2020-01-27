package stockspans

// GetSpans get the spans of the stock prices
func Spans(stocks []int) []int {
	spans := make([]int, len(stocks))
	for i := 0; i < len(stocks); i++ {
		span := 1
		for j := i - 1; j >= 0; j-- {
			if stocks[i] < stocks[j] {
				break
			}
			span++
		}
		spans[i] = span
	}
	return spans
}
