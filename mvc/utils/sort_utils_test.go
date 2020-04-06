package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization
	elements := []int{9, 8, 7, 6}	
	// Execution
	elements = BubbleSort(elements)
	// Validation
	assert.NotNil(t, elements)

	sortedElements := []int{6, 7, 8, 9}
	for idx, val := range sortedElements {
		assert.Equal(t, elements[idx], val)
	}
}


// Bun
func BenchmarkBubbleSort100(b *testing.B) {
	// Initialization
	elements := getElements(100, false)
	// Execution
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	// Initialization
	elements := getElements(1000, false)
	// Execution
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkGoSortInts1000(b *testing.B){
	elements := getElements(1000, false)
	for i := 0; i < b.N; i++ {
		sort.Ints(elements)
	}
}

//
func BenchmarkSortConditional10000(b *testing.B){
	elements := getElements(10000, false)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

// getElements
// Factory method to get elements.
// If `ascending`, returns sorted list of ints
// If not `ascending` returns reversed list of ints
func getElements(n int, ascending bool) []int{
	if ascending {
		return getElementsAsc(n)
	}
	return getElementsDesc(n)
}

// getElementsDesc
// Create []int of values in descending order
// from `n` to 0
func getElementsDesc(n int) []int{
	result := make([]int, n)
	i := 0
	for j := n-1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

// getElementsAsc
// Create []int of values in ascending order
// from 0 to `n`
func getElementsAsc(n int) []int{
	result := make([]int, n)
	for i := 0; i < n; i++{
		result[i] = i
	}
	return result
}
