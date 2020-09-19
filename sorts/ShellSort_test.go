package sorts

import (
	"testing"
)

func ShellSort(a []int, n int) {
	if n <= 1 {
		return
	}
	step := n / 2
	for step > 0 {
		for i := step; i < n; i++ {
			j := i
			for j >= step && a[j] < a[j-step] {
				a[j], a[j-step] = a[j-step], a[j]
				j -= step
			}
		}
		step /= 2
	}

}

func TestShellSort(t *testing.T) {
	a := generateRandonArray(100)
	ShellSort(a, len(a))
	t.Log(a)
}
