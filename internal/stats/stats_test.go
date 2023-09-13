package stats

import (
	"math/rand"
	"testing"
)

func TestPercentiles(t *testing.T) {
	tests := []struct {
		values      []int
		expected    [3]int
		description string
	}{
		// Expected : 15, 20, 25 if 99 nth value is approximated to the 100th one
		{[]int{5, 10, 15, 20, 25}, [3]int{15, 20, 20}, "Percentiles for sorted values"},
		{[]int{1, 1, 1, 1, 1}, [3]int{1, 1, 1}, "Percentiles for identical values"},
		// Expected : 5, 8, 9 if 99 nth value is approximated to the 100th one
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, [3]int{5, 8, 8}, "Percentiles for reverse sorted values"},
	}

	for _, test := range tests {
		result := Percentiles(true, test.values)
		if result != test.expected {
			t.Errorf("%s: PercentilesV2(%v) = %v; expected %v", test.description, test.values, result, test.expected)
		}
	}
}

func BenchmarkPercentiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Percentiles(false, rand.Perm(100))
	}
}
