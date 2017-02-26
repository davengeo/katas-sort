package main

import (
	"fmt"
	"strconv"
	"reflect"
	//"os"
	"os"
	"bufio"
)


/*
	TODO:
1.  Fill the array from lines in a file
2.  Try to define the big O of the current algorithm using different files
3.  Beat the bubble-sort with a better performing algorithm
4.  Unit-test the implementation

*/

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



func addToArray(pos int, value Any, arr []Any) {
	arr[pos] = value
}

func bubbleSort(arr []Any) ([]Any, int) {
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



func main() {

	path := os.Args[1]
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	var theArray [6]Any

	addToArray(0, NewAny(300), theArray[:])
	addToArray(1, NewAny(2), theArray[:])
	addToArray(2, NewAny("b"), theArray[:])
	addToArray(3, NewAny(58), theArray[:])
	addToArray(4, NewAny("a"), theArray[:])
	addToArray(5, NewAny("a23"), theArray[:])

	resultArray, numberOfIter := bubbleSort(theArray[:])

	printArray(resultArray)
	fmt.Println(numberOfIter)
}

func printArray(arr []Any) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i].value)
	}
}