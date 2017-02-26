package impl

import (
	"reflect"
	"strconv"
)

type Any struct {
	value interface{}
}

func NewAny(value interface{}) Any {
	any := new(Any)
	any.value = value
	return *any
}

func (first Any) compare (second Any) int {
	strFirst := first.toString()
	strSecond := second.toString()

	if (strFirst > strSecond) {
		return 1
	}
	if (strFirst < strSecond){
		return -1
	}
	return 0

}

func (v Any) toString() string {
	if (reflect.TypeOf(v.value).Kind() == reflect.Int) {
		return strconv.Itoa(v.value.(int))
	} else {
		return v.value.(string)
	}
}


func BubbleSort(arr []Any) ([]Any, int) {
	var result = make([]Any, len(arr))
	copy(result[:], arr)

	needsMoreIter := true
	numberOfIter := 0

	for needsMoreIter {

		numberOfIter++;
		needsMoreIter = false

		for i := 0; i < (len(result) - 1); i++ {
			if (result[i].compare(result[i+1]) == 1) {
				needsMoreIter = true
				result[i], result[i+1] = result[i+1], result[i]
			}
		}

	}

	return result, numberOfIter
}

func IsSorted(arr []Any) bool {
	for i := 0; i < (len(arr) - 1); i++ {
		if (arr[i].compare(arr[i+1]) == 1) {
			return false
		}
	}
	return true
}
