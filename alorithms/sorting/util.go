package sorting

import "math/rand"

// RandomArray is
func RandomArray(size, rng int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(rng)
	}
	return arr
}
