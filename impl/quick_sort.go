package impl

func QuickSort(arr []string) ([]string, int) {
	doQuicksort(arr, 1, len(arr)-1)
	return arr, -1
}

func doQuicksort(arr []string, lo int, hi int) {
	//fmt.Println("call 1")
	if lo < hi {
		p := calculatePivot(arr, lo, hi)
		doQuicksort(arr, lo, p-1)
		doQuicksort(arr, p+1, hi)
	}
}

func calculatePivot(arr []string, lo int, hi int) int {
	pivot := arr[hi]
	i := lo - 1
	for j := lo; j < hi-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[i+1], arr[hi] = arr[hi], arr[i+1]
	return i + 1
}
