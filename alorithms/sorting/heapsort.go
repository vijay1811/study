package sorting

// Heapsort sorts the array using heap
func Heapsort(arr []int) {
	MinHeap(arr)
	for i := 0; i < len(arr); i++ {
		arr[len(arr)-1-i], arr[0] = arr[0], arr[len(arr)-1-i]
		percolateDown(arr, 0, len(arr)-1-i)
	}
}

// MinHeap creates min heap
func MinHeap(arr []int) {
	HeapifyMin(arr)
}

// HeapifyMin ...
func HeapifyMin(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		percolateDown(arr, i, len(arr))
	}
}

func percolateDown(arr []int, idx, heapLength int) {
	if heapLength-1 < 2*idx+1 {
		// this is a leaf element
		return
	}

	lc := arr[2*idx+1]
	if heapLength-1 < 2*idx+2 {
		if lc > arr[idx] {
			arr[2*idx+1], arr[idx] = arr[idx], arr[2*idx+1]
			percolateDown(arr, 2*idx+1, heapLength)
		}
		return
	}
	rc := arr[2*idx+2]

	if lc > rc {
		if lc > arr[idx] {
			arr[2*idx+1], arr[idx] = arr[idx], arr[2*idx+1]
			percolateDown(arr, 2*idx+1, heapLength)
		}
		return
	}

	if rc > arr[idx] {
		arr[2*idx+2], arr[idx] = arr[idx], arr[2*idx+2]
		percolateDown(arr, 2*idx+2, heapLength)
	}

}
