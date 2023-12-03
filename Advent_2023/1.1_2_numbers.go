package main

import (
	"Advent_2023/day1"
	"Advent_2023/utility"
	"fmt"
	"sync"
	"time"
)

func worker(strings <-chan []byte, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for s := range strings {
		firstCh := make(chan int)
		lastCh := make(chan int)
		day1.FindFirstDigit(s, firstCh)
		day1.FindLastDigit(s, lastCh)

		first, okFirst := <-firstCh
		last, okLast := <-lastCh

		sum := 0
		if okFirst {
			sum += first * 10
		}
		if okLast {
			sum += last
		}

		results <- sum
	}
}

func main() {
	startTime := time.Now()

	iterator, err := utility.LineIteratorBytes("day1_data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	workerCount := 1
	linesCh := make(chan []byte)
	resultsCh := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(linesCh, resultsCh, &wg)
	}
	go func() {
		for {
			line, ok := iterator()
			if !ok {
				break
			}
			linesCh <- line
		}
		close(linesCh)
	}()

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	finalSum := 0
	for sum := range resultsCh {
		finalSum += sum
	}

	duration := time.Since(startTime)
	fmt.Printf("The final sum is: %d\n", finalSum)
	fmt.Printf("Time taken: %s\n", duration)
}
