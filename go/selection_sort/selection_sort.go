package main

import "fmt"

func pop(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for i, _ := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selectionSort(arr []int) []int {
	newArr := make([]int, len(arr))
	for i, _ := range arr {
		smallest := findSmallest(arr)
		newArr[i] = arr[smallest]
		arr = pop(arr, smallest)
	}
	return newArr
}

func main() {
	fmt.Println(selectionSort([]int{5, 3, 6, 2, 10}))
}
