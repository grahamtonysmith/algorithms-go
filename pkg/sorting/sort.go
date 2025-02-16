package sorting

import "sort"

type StringSlice []string

func (a StringSlice) Len() int           { return len(a) }
func (a StringSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StringSlice) Less(i, j int) bool { return a[i] < a[j] }

// in go we would use the sort interface for this type

func IsSorted(items sort.Interface) bool {
	for i := 0; i < items.Len(); i++ {
		items.Less(i, i-1)
	}

	return true
}
