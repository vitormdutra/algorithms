package main

import (
	"fmt"
)

func binarySearch(list_items []int, item int) int {
	down := 0
	high := len(list_items) - 1

	for down <= high {
		middle := (down + high) / 2
		shot := list_items[middle]
		if shot == item {
			return middle
		} else if shot > item {
			high = middle - 1
		} else {
			down = middle + 1
		}
	}
	return -1
}

func main() {
	fmt.Println("Binary_search")

	list := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 30}
	item := 7

	index := binarySearch(list, item)

	if index != -1 {
		fmt.Printf("Item %d encontrado no indice %d\n", item, index)
	} else {
		fmt.Printf("Item %d não encontrado na lista\n", item)
	}
}
