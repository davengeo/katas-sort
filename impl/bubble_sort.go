package impl

func BubbleSort(arr []string) ([]string, int) {
	var result = make([]string, len(arr))
	copy(result[:], arr)

	needsMoreIter := true
	numberOfIter := 0

	for needsMoreIter {

		numberOfIter++;
		needsMoreIter = false

		for i := 0; i < (len(result) - 1); i++ {
			if (result[i] > result[i+1]) {
				needsMoreIter = true
				result[i], result[i+1] = result[i+1], result[i]
			}
		}

	}

	return result, numberOfIter
}

func IsSorted(arr []string) bool {
	for i := 0; i < (len(arr) - 1); i++ {
		if (arr[i] > arr[i+1]) {
			return false
		}
	}
	return true
}
