package impl

import "fmt"

func QuickSort(arr []string) ([]string, int) {
	if isSorted(arr) {
		return arr, -1
	}
	pvt:=calculatePivot(arr)
	fmt.Println(arr[pvt])
	left, right:=getLeftRight(arr, pvt)
	return QuickSort(append(left, right...))
}

func calculatePivot(arr []string) (int) {
	return (len(arr)/2 -1)
}


func getLeftRight(arr []string, pvt int) (left []string, right []string) {
	left=make([]string, 0)
	right=make([]string, 0)

	for i:=0; i<len(arr); i++ {
		if(arr[i] > arr[pvt]) {
			right=append(right, arr[i])
		} else {
			left=append(left, arr[i])
		}
	}
	return left, right
}

func isSorted(arr []string) bool {
	for i := 0; i < (len(arr) - 1); i++ {
		if (arr[i] > arr[i+1]) {
			return false
		}
	}
	return true
}
