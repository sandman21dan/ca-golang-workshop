package demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func fakeHttpRequest(url string) string {
	time.Sleep(100 * time.Millisecond)

	return "response"
}

func TestConcurrency(t *testing.T) {
	t.Run("Fixed concurrency", func(t *testing.T) {
		start := time.Now()

		responses := make(chan string, 2)

		go func() {
			responses <- fakeHttpRequest("url1")
		}()

		go func() {
			responses <- fakeHttpRequest("url2")
		}()

		// Can perform calculations in the mean time
		time.Sleep(50 * time.Millisecond)

		r1 := <-responses
		r2 := <-responses

		fmt.Println("Response from url1", r1)
		fmt.Println("Response from url2", r2)

		fmt.Printf("Took %s\n", time.Since(start))
	})

	t.Run("Variable concurrency", func(t *testing.T) {
		start := time.Now()
		someExpensiveCalculation := func(v int) int {
			time.Sleep(20 * time.Millisecond)
			return v * v
		}

		values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		concurrency := 6
		valChan := make(chan int)
		resChan := make(chan int)
		wg := &sync.WaitGroup{}

		// start goroutines for the desired concurrency
		for i := 0; i < concurrency; i++ {
			wg.Add(1)

			go func() {
				for v := range valChan {
					fmt.Println("Processing", v)
					resChan <- someExpensiveCalculation(v)
				}

				wg.Done()
			}()
		}

		// read results as they get calculated
		go func() {
			for v := range resChan {
				fmt.Println("Result", v)
			}
		}()

		// send values to be calculated
		for _, v := range values {
			valChan <- v
		}
		// close sending channel to signal end of values
		close(valChan)

		// wait for all goroutines to finish
		wg.Wait()
		close(resChan)

		fmt.Printf("Took %s\n", time.Since(start))
	})
}
