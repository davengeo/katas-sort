package impl

import "fmt"

func printArray(arr []string) {
	fmt.Println("*********")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
