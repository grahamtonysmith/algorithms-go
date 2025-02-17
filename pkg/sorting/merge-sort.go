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

func mergeSort(items StringSlice, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + ((hi - lo) / 2)

	mergeSort(items, lo, mid)   // sort left
	mergeSort(items, mid+1, hi) // sort right

	merge(items, lo, mid, hi) // merge results
}

func MergeSort(items StringSlice) {
	lo := 0
	hi := len(items) - 1

	mergeSort(items, lo, hi)
}
