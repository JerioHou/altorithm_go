package sorts

import (
	"testing"
)

func selectionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
}

func TestSelectionSort(t *testing.T) {
	a := []int{12, 23, 4, 51, 23, 4, 9, 78, 35, 49, 13, 97, 1235, 4}
	selectionSort(a, len(a))
	t.Log(a)

}
