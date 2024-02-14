package main

import (
	"log"
	"github.com/jackelder/algorithms/mergesort"
)

func main() {
	log.Println("merge sort demo")

	input := []int{23, 2, 87, 19}

	log.Println("Unordered array:", input)

	output := mergesort.MergeSort(input)

	log.Println("Sorted array:", output)
}

