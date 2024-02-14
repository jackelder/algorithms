package mergesort

import (
	"math"
)

// MergeSort an array of integers
func MergeSort(input []int) []int {
	return mergeSort(input, len(input))
}

// Use the MergeSort algorithm to sort an array of integers (input) with length n
func mergeSort(input []int, n int) []int {
	if n <= 1 {
		return input
	}

	halfwayFloat := math.Ceil(float64(n) / 2)
	halfway := int(halfwayFloat)
	first_half := input[:halfway]
	second_half := input[halfway:]

	b := mergeSort(first_half, len(first_half))
	c := mergeSort(second_half, len(second_half))
	d := merge(b, c, n)
	return d
}

func merge(a, b []int, n int) []int {
	aLength := len(a) // use better length accounting to remove this check
	bLength := len(b)
	c := make([]int, n)
	i := 0
	j := 0

	for k := 0; k < n; k++ {
		if i >= aLength && j >= bLength {
			break
		}

		if j >= bLength {
			c[k] = a[i]
			i += 1
		} else if i >= aLength {
			c[k] = b[j]
			j += 1
		} else if a[i] <= b[j] {
			c[k] = a[i]
			i += 1
		} else {
			c[k] = b[j]
			j += 1
		}
	}

	return c
}
