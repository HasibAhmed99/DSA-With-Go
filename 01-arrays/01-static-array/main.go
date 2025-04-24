package main

import (
	"fmt"
)

// insertEnd inserts the value `n` into the next open position in `arr`
// if the current logical length is less than the array's capacity.
// Note: Length is the logical size, capacity is the total allocated size.
func insertEnd(arr []int, n int, length *int, capacity int) {
	if *length < capacity {
		arr[*length] = n
		*length++ // Increase logical length
	}
}

// removeEnd removes the last value in the array by overwriting it
// with a default zero and decreasing the logical length.
// It does not shrink the underlying slice.
func removeEnd(arr []int, length *int) {
	if *length > 0 {
		*length--
		arr[*length] = 0 // Optional: zero out the removed value
	}
}

// insertMiddle inserts value `n` at index `i` by shifting elements to the right.
// It assumes that index `i` is valid and the array is not already full.
func insertMiddle(arr []int, i int, n int, length *int, capacity int) {
	if *length >= capacity || i < 0 || i > *length {
		fmt.Println("Error: Insertion out of bounds or array full.")
		return
	}

	// Shift elements from the end to index i
	for index := *length - 1; index >= i; index-- {
		arr[index+1] = arr[index]
	}

	arr[i] = n
	*length++
}

// removeMiddle removes the element at index `i` by shifting elements
// to the left starting from index i+1. The logical length is reduced.
func removeMiddle(arr []int, i int, length *int) {
	if i < 0 || i >= *length {
		fmt.Println("Error: Removal out of bounds.")
		return
	}

	for index := i + 1; index < *length; index++ {
		arr[index-1] = arr[index]
	}

	*length--
	arr[*length] = 0 // Optional: clear trailing value
}

// printArr prints all elements of the array up to the given capacity.
// This simulates printing the full backing array, not just logical values.
func printArr(arr []int, capacity int) {
	for i := 0; i < capacity; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func main() {
	capacity := 10
	arr := make([]int, capacity)
	length := 0

	insertEnd(arr, 10, &length, capacity)
	fmt.Println(length)
	insertEnd(arr, 20, &length, capacity)
	insertEnd(arr, 30, &length, capacity)
	printArr(arr, capacity) // Output: 10 20 30 0 0 0 0 0 0 0

	insertMiddle(arr, 1, 15, &length, capacity)
	printArr(arr, capacity) // Output: 10 15 20 30 0 0 0 0 0 0

	removeMiddle(arr, 2, &length)
	printArr(arr, capacity) // Output: 10 15 30 0 0 0 0 0 0 0

	removeEnd(arr, &length)
	printArr(arr, capacity) // Output: 10 15 0 0 0 0 0 0 0 0
}
