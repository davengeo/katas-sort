package impl

func MergeSort(arr []string) ([]string, int) {
	return doMergeSort(arr), -1
}

// Devide part of the algo
func doMergeSort(arr []string) []string {
	length := len(arr)

	if length <= 1 {
		return arr
	}

	pivot := int(length / 2)
	right := arr[:pivot]
	left := arr[pivot:]

	return merge(doMergeSort(left), doMergeSort(right))
}

// Conquer part of the algo
func merge(left, right []string) []string {

	lIndex, rIndex := 0, 0
	resultLength := len(left) + len(right)
	result := make([]string, resultLength)

	for i := 0; i < resultLength; i++ {
		if lIndex > len(left)-1 && rIndex <= len(right)-1 {
			result[i] = right[rIndex]
			rIndex++
		} else if rIndex > len(right)-1 && lIndex <= len(left)-1 {
			result[i] = left[lIndex]
			lIndex++
		} else if left[lIndex] < right[rIndex] {
			result[i] = left[lIndex]
			lIndex++
		} else {
			result[i] = right[rIndex]
			rIndex++
		}
	}

	return result
}
