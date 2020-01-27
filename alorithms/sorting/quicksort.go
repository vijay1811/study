package sorting

// Quicksort ...
func Quicksort(arr []int) {
	if len(arr) == 0 {
		return
	}
	if len(arr) == 1 {
		return
	}
	pivot := partition(arr)
	Quicksort(arr[:pivot])
	if pivot+1 == len(arr) {
		return
	}
	Quicksort(arr[pivot+1:])

}

func partition(arr []int) int {
	pivot := arr[0]
	i := 1
	j := len(arr) - 1
	for j >= i {
		if arr[i] < pivot {
			i++
			continue
		}
		if arr[j] > pivot {
			j--
			continue
		}
		arr[j], arr[i] = arr[i], arr[j]
		i++
		j--
	}
	arr[0], arr[j] = arr[j], arr[0]
	return j
}
