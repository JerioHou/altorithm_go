package sorts

import (
	"testing"
)

func MergeSort(a []int, result []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	start1 := start
	start2 := mid + 1
	MergeSort(a, result, start1, mid)
	MergeSort(a, result, start2, end)
	k := start
	for ; start1 <= mid && start2 <= end; k++ {
		if a[start1] <= a[start2] {
			result[k] = a[start1]
			start1++
		} else {
			result[k] = a[start2]
			start2++
		}
	}
	for ; start1 <= mid; start1++ {
		result[k] = a[start1]
		k++
	}
	for ; start2 <= end; start2++ {
		result[k] = a[start2]
		k++
	}
	copy(a[start:end+1], result[start:end+1])
}

func TestMergeSort(t *testing.T) {
	a := generateRandonArray(7)
	t.Log(a)
	result := make([]int, len(a))
	MergeSort(a, result, 0, len(a)-1)
	t.Log(a)
}
