package sorts

import "testing"

func quickSort(a []int, l int, r int) {
	if l >= r {
		return
	}
	idx := paritition(a, l, r)
	quickSort(a, l, idx-1)
	quickSort(a, idx+1, r)
}

func paritition(a []int, l int, r int) int {
	// 基准住,取最后一位
	pivot := a[r]
	// 游标，标记当前小于基准值得 下一个下标
	idx := l
	for i := l; i < r; i++ {
		if a[i] < pivot {
			if i == idx {
				idx++
			} else {
				a[idx], a[i] = a[i], a[idx]
				idx++
			}
		}
	}
	a[idx], a[r] = a[r], a[idx]
	return idx
}

func TestIQuickSort(t *testing.T) {
	a := generateRandonArray(100)
	quickSort(a, 0, len(a)-1)
	t.Log(a)
}
