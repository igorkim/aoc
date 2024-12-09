package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := flag.String("f", "input.txt", "input file")
	flag.Parse()

	arr1, arr2, err := readInput(*filename)
	if err != nil {
		log.Fatalf("readInput error - %s", err)
	}

	dist := calculateDistance(arr1, arr2)
	log.Printf("DISTANCE: %d", dist)

	sim := calculateLeftSimilarity(arr1, arr2)
	log.Printf("SIMILARITY: %d", sim)
}

func readInput(filename string) ([]uint, []uint, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var arr1, arr2 []uint
	for scanner.Scan() {
		text := scanner.Text()

		recs := strings.Split(text, "   ")
		if len(recs) != 2 {
			return nil, nil, fmt.Errorf("invalid number of arguments - %q", text)
		}

		a1, err := strconv.ParseUint(recs[0], 10, 64)
		if err != nil {
			return nil, nil, err
		}
		arr1 = append(arr1, uint(a1))

		a2, err := strconv.ParseUint(recs[1], 10, 64)
		if err != nil {
			return nil, nil, err
		}
		arr2 = append(arr2, uint(a2))

	}

	if err := scanner.Err(); err != nil {
		// return nil, nil, err
	}

	return arr1, arr2, nil
}

// arrays must be sorted
func calculateLeftSimilarity(arr1, arr2 []uint) uint {
	var sim uint
	var lastSim uint
	var jIdx int
	for i := 0; i < len(arr1); i++ {
		if i > 0 && arr1[i] == arr1[i-1] {
			sim += lastSim
			continue
		}

		var count uint
		for j := jIdx; j < len(arr2); j++ {
			if arr1[i] < arr2[j] {
				lastSim = arr1[i] * count
				sim += lastSim
				jIdx = j
				break
			}

			// skip
			if arr1[i] > arr2[j] {
				continue
			}

			// arr1[i] == arr2[j]
			count++
		}
	}

	return sim
}

func calculateDistance(arr1, arr2 []uint) uint {
	bucketSort(arr1)
	bucketSort(arr2)

	var dist uint
	for i := 0; i < len(arr1); i++ {
		if arr1[i] > arr2[i] {
			dist += arr1[i] - arr2[i]
			continue
		}
		dist += arr2[i] - arr1[i]
	}

	return dist
}

func bucketSort(arr []uint) {
	min, max := findMinMax(arr)

	buckets := make([]uint, max-min+1)

	for i := 0; i < len(arr); i++ {
		idx := arr[i] - min
		buckets[idx]++
	}

	var idx int
	for i := 0; i < len(buckets); i++ {
		for j := uint(0); j < buckets[i]; j++ {
			arr[idx] = min + uint(i)
			idx++
		}
	}
}

func findMinMax(arr []uint) (uint, uint) {
	var min uint = 1 << 63
	var max uint

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	return min, max
}
