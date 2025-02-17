package sorting

// merge creates a copy of items, aux, and two pointers i and j pointing initially to low and hi
// it then iters over each item in items with index k, and compares the values of the points i and j
// if i is > then mid, then we take from the right
// if j > high, then we take from the left
// if value on the right is less than the value on the left, then take from right, move j right one
// else the value on the right is greater than or equal to the value of left, take from the left
// taking from means writing the values into the items array
func merge(items StringSlice, lo, mid, hi int) {
	// pretty bad that we do this on every merge call, we should make a struct tbh
	aux := make(StringSlice, len(items))
	copy(aux, items)

	i, j := lo, mid+1

	// for each item in lo to hi, i.e., for each item in the array
	for k := lo; k <= hi; k++ {
		if i > mid {
			// left hand exhausted, take from right
			items[k] = aux[j]
			j++
		} else if j > hi {
			// right hand exhausted, take from left
			items[k] = aux[i]
			i++
		} else if aux.Less(j, i) {
			// current key on right less than current key on left, take from right
			items[k] = aux[j]
			j++
		} else {
			// current key on right greater than or equal to current key on left, take from left
			items[k] = aux[i]
			i++
		}
	}
}

// we can make this faster in 3 ways
// using insertion sort for small arrays (because resursion meads we'll get to this case)
// testing if the array is in order already
// use resursive trickery to eliminate the time take to copy the aux array
func mergeSort(items StringSlice, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + ((hi - lo) / 2)

	mergeSort(items, lo, mid)   // sort left
	mergeSort(items, mid+1, hi) // sort right

	merge(items, lo, mid, hi) // merge results
}

// great example of devide and conquer strategy
// break a big problem into smaller ones and merge results
func MergeSort(items StringSlice) {
	lo := 0
	hi := len(items) - 1

	mergeSort(items, lo, hi)
}

// build small solutions into larger ones
func MergeSortBU(items StringSlice) {
	n := len(items)

	for sz := 1; sz < n; sz += sz {
		for lo := 0; lo < n-sz; lo += sz + sz {
			mid := lo + sz - 1
			hi := min(lo+sz+sz-1, n-1)

			merge(items, lo, mid, hi)
		}
	}
}
