package sorts

import "testing"

// 插入排序
func InsertSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		tmp := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if tmp < a[j] {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = tmp
	}
}

// 测试插入排序
func TestInsertSort(t *testing.T) {
	a := generateRandonArray(100)
	InsertSort(a, len(a))
	t.Log(a)
}
