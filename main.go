package main

import (
	"./impl"
	"./report"
	"bufio"
	"fmt"
	"os"
	"time"
)

func loadFile(path string) (lines []string, numberLines int) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	numberLines = 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
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

	resultBubble, _ := impl.BubbleSort(theArray[:])

	printInterval("bubbleSort", t)
	fmt.Println("sorted:", impl.IsSorted(resultBubble))

	r.AddItem("bubble", os.Args[1], lines, time.Now().Sub(t).Nanoseconds(), impl.IsSorted(resultBubble))

	t = time.Now()
	resultTree, _ := impl.SortTree(theArray[:])

	printInterval("treeSort", t)
	fmt.Println("sorted:", impl.IsSorted(resultTree))

	r.AddItem("tree", os.Args[1], lines, time.Now().Sub(t).Nanoseconds(), impl.IsSorted(resultTree))

	t = time.Now()
	resultQuick, _ := impl.QuickSort(theArray[:])

	printInterval("quickSort", t)
	fmt.Println("sorted:", impl.IsSorted(resultQuick))

	r.AddItem("quick", os.Args[1], lines, time.Now().Sub(t).Nanoseconds(), impl.IsSorted(resultQuick))

	t = time.Now()
	resultMerge, _ := impl.MergeSort(theArray[:])

	printInterval("mergeSort", t)
	fmt.Println("sorted:", impl.IsSorted(resultMerge))

	r.AddItem("merge", os.Args[1], lines, time.Now().Sub(t).Nanoseconds(), impl.IsSorted(resultMerge))

	r.WriteToFile()
}
