package stats

import "math/rand"

// Pivot choose a random value to partition the v list.
func Pivot(v []int) int {
	n := len(v)
	return Median(v[rand.Intn(n)],
		v[rand.Intn(n)],
		v[rand.Intn(n)])
}

// Median return the median value of three values.
func Median(a, b, c int) int {
	if a < b {
		switch {
		case b < c:
			return b
		case a < c:
			return c
		default:
			return a
		}
	}

	switch {
	case a < c:
		return a
	case b < c:
		return c
	default:
		return b
	}
}

// Partition reorders the elements of list so that:
// - all elements in list[:low] are less than p,
// - all elements in list[low:high] are equal to p,
// - all elements in list[high:] are greater than p.
func Partition(list []int, p int) (low, high int) {
	low, high = 0, len(list)
	for mid := 0; mid < high; {
		// Loop Invariant:
		//  - v[:low] < p
		//  - v[low:mid] = p
		//  - v[mid:high] are unknown
		//  - v[high:] > p
		switch a := list[mid]; {
		case a < p:
			list[mid] = list[low]
			list[low] = a
			low++
			mid++
		case a == p:
			mid++
		default: // a > p
			list[mid] = list[high-1]
			list[high-1] = a
			high--
		}
	}

	return
}

// InsertionSort is a simple sorting algorithm that builds the final
// sorted array (or list) one item at a time by comparisons.
func InsertionSort(v []int) {
	for j := 1; j < len(v); j++ {
		// Invariant: v[:j] contains the same elements as
		// the original slice v[:j], but in sorted order.
		key := v[j]

		i := j - 1
		for i >= 0 && v[i] > key {
			v[i+1] = v[i]
			i--
		}

		v[i+1] = key
	}
}

// Quicksort sorts the elements of v in ascending order using InsertionSort for small (length < 20) list
// and PartitionSort for bigger ones.
func Quicksort(v []int) {
	if len(v) < 20 {
		InsertionSort(v)
		return
	}

	p := Pivot(v)

	low, high := Partition(v, p)
	Quicksort(v[:low])
	Quicksort(v[high:])
}
