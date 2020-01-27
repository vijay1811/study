package sorting

// InsertionSort ...
func InsertionSort(arr []int) {
	if len(arr) == 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		e := arr[i]
		j := i - 1
		for ; j >= 0 && e < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = e
	}
}

//
