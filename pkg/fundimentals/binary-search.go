package fundimentals

import "fmt"

func BinarySrarch(arr []int, key int) int {
	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := lo + ((hi - lo) / 2)

		fmt.Println("lo", lo, "hi", hi, "mid", mid)

		if key < arr[mid] {
			hi = mid - 1
		} else if key > arr[mid] {
			lo = mid + 1
		} else {
			return mid
		}

		fmt.Println("lo", lo, "hi", hi, "mid", mid)
	}

	return -1
}
