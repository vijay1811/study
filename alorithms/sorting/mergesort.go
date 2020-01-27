package sorting

// Mergesort sorts array
func Mergesort(arr []int) {
	ms(arr, 0, len(arr))
}

func ms(arr []int, start, end int) {
	if start >= end-1 {
		return
	}
	mid := (start + end) / 2
	ms(arr, start, mid)
	ms(arr, mid, end)
	Merge(arr, start, mid, end)
}

// Merge merges two sorted array
func Merge(arr []int, start, mid, end int) {
	lowerArr := arr[start:mid]
	upperArr := arr[mid:end]
	mArr := make([]int, end-start)
	i, j, k := 0, 0, 0
	for ; i < len(lowerArr) && j < len(upperArr); k++ {
		if lowerArr[i] < upperArr[j] {
			mArr[k] = lowerArr[i]
			i++
			continue
		}
		mArr[k] = upperArr[j]
		j++
	}
	if i == len(lowerArr) && j < len(upperArr) {
		mArr = append(mArr[:k], upperArr[j:]...)
	}
	if j == len(upperArr) && i < len(lowerArr) {
		mArr = append(mArr[:k], lowerArr[i:]...)
	}

	for i := 0; i < end-start; i++ {
		arr[i+start] = mArr[i]
	}
}
