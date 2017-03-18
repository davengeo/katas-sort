package main

import (
	"fmt"
	"os"
	"bufio"
	"time"
	"./impl"
	"./report"
)


/*
	TODO:
1.  Fill the array from lines in a file. DONE
2.  Try to define the big O of the current algorithm using different files
3.  Beat the bubble-sort with a better performing algorithm

*/
func loadFile(path string ) (lines []impl.Any, numberLines int) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	numberLines = 0

	for scanner.Scan() {
		lines = append(lines, impl.NewAny(scanner.Text()))
		numberLines++
	}
	return
}



func printInterval(title string, t time.Time) {
	fmt.Printf("%s:%s \n", title, time.Now().Sub(t).String())
}

func main() {
	r := report.NewReport("dummy", "dummy.csv")
	t := time.Now()


	theArray, lines := loadFile(os.Args[1])
	printInterval("loadFile", t)
	fmt.Println("lines:", lines)

	t = time.Now()
	//result, numberOfIter := impl.BubbleSort(theArray[:])
	result, numberOfIter := impl.TreeSort(theArray[:])
	printInterval("treeSort", t)

	//printArray(resultArray)

	fmt.Println("iters:",numberOfIter)

	fmt.Println("sorted:",impl.IsSorted(result))

	r.AddItem("bubble", os.Args[1], lines, time.Now().Sub(t).Nanoseconds(), impl.IsSorted(result))
	r.WriteToFile()
}




