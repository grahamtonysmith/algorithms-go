package sorting

import (
	"sort"
)

// insertion sort is "slow" because only adjacent entries can be exchanged
// for example, if the smallest entry is at the end of an array then N-1 exchanges are required to get it where it belongs
// shellsort gains speed by allowing entries that are far away to be exchanged, creating a partially sorted array to be finished off by selection sort
func ShellSort(a sort.Interface) {
	n := a.Len()
	h := 1

	for h < n/3 {
		h = (3 * h) + 1 // 1, 4, 13, 40, ...
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && a.Less(j, j-h); j -= h {
				a.Swap(j, j-h)
			}
		}

		h = h / 3
	}
}
