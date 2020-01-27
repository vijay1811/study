package sorting

// SelectionSort ...
func SelectionSort(arr []int) {
	for i := range arr {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}
