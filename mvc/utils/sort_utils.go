package utils

import "sort"

// BubbleSort
// A simple yet inefficient sorting algorithm
func BubbleSort(elements []int) []int{
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements) - 1; i ++ {
			if elements[i] > elements [i+1]	{
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}

	}
	return elements
}

func Sort(elements []int) []int{
	if len(elements) < 1000 {
		return BubbleSort(elements)
	}
	sort.Ints(elements) // Inplace
	return elements
}
