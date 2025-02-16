package sorting

import "sort"

// nearly linear for nearly sorted arrays
// so great for non random data
// arrays where each entry is not far from its final position
// arrays where only a few entries are not in place
// arrays where a smaller array is appended to a larger sorted array
// can be sped up substantially by moving larger values to the right by one position rather than doing full exchanges
func InsertionSort(a sort.Interface) {
	n := a.Len()

	for i := 1; i < n; i++ {
		// insert a[i] among a[i-1], a[i-2], ...

		for j := i; j > 0 && a.Less(j, j-1); j-- {
			a.Swap(j, j-1)
		}
	}
}
