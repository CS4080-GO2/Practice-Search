package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test := generateRandomSortedArray(100000)
	fmt.Printf("Array is sorted?: %v\n", verify(test))

	fmt.Print("Binary Search:\n")
	start := time.Now()
	result := fullTest(test)
	end := time.Now()
	fmt.Printf("worked?: %v. Took: %d ms\n", result, (end.Nanosecond()/1000000 - start.Nanosecond()/1000000))

	fmt.Print("Exponential Search:\n")

	newStart := time.Now()
	result = fullExpoTest(test)
	newEnd := time.Now()
	fmt.Printf("worked?: %v. Took: %d ms", result, (newEnd.Nanosecond()/1000000 - newStart.Nanosecond()/1000000))
}

func generateRandomSortedArray(size int) []int {
	test := make([]int, size) // Had to make slice instead of array. Almost the same.
	prev := 0
	for i := 0; i < size; i++ {
		prev = rand.Intn(5) + 1 + prev
		test[i] = prev
	}
	return test
}

func verify(arr []int) bool {
	for i, v := range arr {
		if i == 0 {
			continue
		}
		if !(arr[i-1] < v) {
			return false
		}
	}
	return true
}

func fullTest(arr []int) bool {
	size := len(arr)
	for i, v := range arr {
		//fmt.Println(i)
		if binarySearch(arr, v, 0, size) != i {
			return false
		}
	}
	return true
}

func fullExpoTest(arr []int) bool {
	for i, v := range arr {
		if expoSearch(arr, v) != i {
			return false
		}
	}
	return true
}

func binarySearch(arr []int, element int, min int, max int) int {
	found := false
	for !found {
		middle := ((max - min) / 2) + min
		eval := arr[middle] - element
		if eval == 0 {
			// Found element
			found = true
			return middle
		} else if eval > 0 {
			// arr[max/2] > element, so go lower in array
			max = middle - 1
		} else {
			// arr[max/2] < element, so go higher in array
			min = middle + 1
		}
	}
	return -1
}

func expoSearch(arr []int, element int) int {
	//return the first index if it's the target element
	if arr[0] == element {
		return 0
	}

	i := 1
	for i < len(arr) && arr[i] <= element {
		i = i * 2
	}

	min := i / 2
	max := i
	//ensure that i <= len(arr)
	if len(arr) < i {
		max = len(arr) - 1
	}

	return binarySearch(arr, element, min, max)

}
