package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSortUnsortedArray(t *testing.T) {
	input := []int{5, 87, 1, 3}
	got := MergeSort(input)
	want := []int{1, 3, 5, 87}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeSort(%v) = %v; want %v", input, got, want)
	}
}

func TestMergeSortDuplicates(t *testing.T) {
	input := []int{5, 3, 87, 3, 1, 3}
	got := MergeSort(input)
	want := []int{1, 3, 3, 3, 5, 87}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeSort(%v) = %v; want %v", input, got, want)
	}
}

func TestMergeSortSortedArray(t *testing.T) {
	input := []int{1, 3, 5, 87}
	got := MergeSort(input)

	if !reflect.DeepEqual(got, input) {
		t.Errorf("MergeSort(%v) = %v; want %v", input, got, input)
	}
}

func TestMergeSortLength1(t *testing.T) {
	input := []int{2}
	got := MergeSort(input)

	if !reflect.DeepEqual(got, input) {
		t.Errorf("MergeSort(%v) = %v; want %v", input, got, input)
	}
}

func TestMergeSortLength0(t *testing.T) {
	input := []int{}
	got := MergeSort(input)

	if !reflect.DeepEqual(got, input) {
		t.Errorf("MergeSort(%v) = %v; want %v", input, got, input)
	}
}
