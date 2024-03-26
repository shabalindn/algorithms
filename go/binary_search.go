package main

import "fmt"

func binary_search(list []int, item int) *int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return &mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nil
}

func main() {
	my_list := []int{1, 3, 5, 7, 9}
	fmt.Println(*binary_search(my_list, 3)) // => 1
	fmt.Println(binary_search(my_list, -1)) // => nil
}
