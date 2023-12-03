package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func findNumber(idx int, greeting string, start int, end int, results *sync.Map) {
	hasher := md5.New()
	prefix := []byte(greeting)
	maxLen := len(greeting) + len(strconv.Itoa(end))
	data := make([]byte, maxLen) // Preallocate memory for data

	copy(data, prefix)

	for number := start; number < end; number++ {
		dataForHash := strconv.AppendInt(data[:len(greeting)], int64(number), 10)
		hasher.Write(dataForHash)
		hash := hasher.Sum(nil)
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			// Store the result in the sync.Map
			results.Store(idx, number)
			break
		}
		hasher.Reset()
	}
}

func main() {
	greeting := "iwrupvqb"
	step := 100000000

	var wg sync.WaitGroup
	results := sync.Map{}

	startTime := time.Now()

	for i := 0; i < 4; i++ {
		wg.Add(1)
		start := i * step
		end := (i + 1) * step
		go func(idx, start, end int) {
			defer wg.Done()
			findNumber(idx, greeting, start, end, &results)
		}(i, start, end)
	}

	wg.Wait() // Wait for all goroutines to finish

	// Print the results stored in the sync.Map
	fmt.Println("Results from goroutines:")
	results.Range(func(key, value interface{}) bool {
		fmt.Printf("Goroutine %d: %d\n", key.(int), value.(int))
		return true
	})

	// Find the minimum number from the results stored in the sync.Map
	var minNumber int = 1<<31 - 1 // Start with maximum int value
	results.Range(func(_, value interface{}) bool {
		if value.(int) < minNumber {
			minNumber = value.(int)
		}
		return true
	})

	elapsedTime := time.Since(startTime)

	fmt.Printf("\nTime taken: %v, Minimum Result: %d\n", elapsedTime, minNumber)
}
