package main

import (
	"fmt"
	"strconv"
	"reflect"
	"os"
	"bufio"
	"time"
)


/*
	TODO:
1.  Fill the array from lines in a file. DONE
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

func fillWithFile(path string ) (lines []Any) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, NewAny(scanner.Text()))
	}
	return
}

func printInterval(t time.Time) {
	fmt.Println(time.Now().Sub(t).String())
}

func main() {

	t := time.Now()
	theArray := fillWithFile(os.Args[1])
	printInterval(t)

	t = time.Now()
	_, numberOfIter := bubbleSort(theArray[:])
	printInterval(t)

	//printArray(resultArray)

	fmt.Println(numberOfIter)
}

func printArray(arr []Any) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i].value)
	}
}