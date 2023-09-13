package stats

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestPivot(t *testing.T) {
	// Create a random input slice
	v := []int{4, 2, 8, 5, 1, 9}
	randomPivot := Pivot(v)

	// Ensure that the pivot is one of the elements in the slice
	found := false

	for _, element := range v {
		if element == randomPivot {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Pivot did not return a valid element from the input slice")
	}
}

func TestMedian(t *testing.T) {
	// Test cases for the Median function
	tests := []struct {
		a, b, c, expected int
	}{
		{3, 2, 1, 2},
		{5, 9, 7, 7},
		{1, 1, 1, 1},
		{7, 4, 9, 7},
		{2, 8, 4, 4},
	}

	for _, test := range tests {
		result := Median(test.a, test.b, test.c)
		if result != test.expected {
			t.Errorf("Median(%d, %d, %d) = %d; expected %d", test.a, test.b, test.c, result, test.expected)
		}
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		inputList                 []int
		p                         int
		expectedLow, expectedHigh int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2, 3},
		{[]int{5, 4, 3, 2, 1}, 3, 2, 3},
		{[]int{1, 1, 1, 1, 1}, 1, 0, 5},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			low, high := Partition(test.inputList, test.p)

			if low != test.expectedLow {
				t.Errorf("Low value is incorrect. Expected: %d, Got: %d", test.expectedLow, low)
			}

			if high != test.expectedHigh {
				t.Errorf("High value is incorrect. Expected: %d, Got: %d", test.expectedHigh, high)
			}

			// Check the partitioned list
			expectedList := make([]int, len(test.inputList))
			copy(expectedList, test.inputList)
			expectedList = expectedList[:low]
			expectedList = append(expectedList, make([]int, test.expectedHigh-low)...)
			for i := low; i < test.expectedHigh; i++ {
				expectedList[i] = test.p
			}
			expectedList = append(expectedList, test.inputList[high:]...)

			if !reflect.DeepEqual(test.inputList, expectedList) {
				t.Errorf("Partitioned list is incorrect. Expected: %v, Got: %v", expectedList, test.inputList)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 2, 8, 5, 1, 9}, []int{1, 2, 4, 5, 8, 9}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{5, 5, 5, 5}, []int{5, 5, 5, 5}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		InsertionSort(test.input)

		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("InsertionSort did not sort the input correctly: %v, expected %v", test.input, test.expected)
		}
	}
}

func TestQuicksort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 2, 8, 5, 1, 9}, []int{1, 2, 4, 5, 8, 9}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{5, 5, 5, 5}, []int{5, 5, 5, 5}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		Quicksort(test.input)

		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("Quicksort did not sort the input correctly: %v, expected %v", test.input, test.expected)
		}
	}
}

func BenchmarkQuicksort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Quicksort(rand.Perm(100))
	}
}
