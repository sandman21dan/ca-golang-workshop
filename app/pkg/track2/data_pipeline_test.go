package pipeline

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

func processPipeline() {
	fileName := "./data.csv.gz"
	log.Println("Reading from file", fileName)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal("Error decoding gzip", err)
	}
	csvr := csv.NewReader(gr)

	concurrency := 10
	bufferSize := 1_000

	wg := &sync.WaitGroup{}
	inChan := make(chan [][2]int)
	outChan := make(chan float64)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(wid int) {
			log.Println("Starting worker", wid)
			for batch := range inChan {
				size := len(batch)
				acum := 0
				for _, row := range batch {
					res := (row[0] * row[1])
					res = res * res
					acum += res
				}
				outChan <- float64(acum) / float64(size)
			}
			wg.Done()
			fmt.Println("Closing worker", wid)
		}(i)
	}

	go func() {
		batches := 0
		buffer := [][2]int{}
		for {
			row, err := csvr.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal("Error reading csv", err)
			}

			col1, err := strconv.Atoi(row[0])
			if err != nil {
				log.Fatal("Error converting col1 to int", err)
			}
			col2, err := strconv.Atoi(row[1])
			if err != nil {
				log.Fatal("Error converting col1 to int", err)
			}

			buffer = append(buffer, [2]int{col1, col2})

			if len(buffer) == bufferSize {
				inChan <- buffer
				batches++
				buffer = [][2]int{}
			}
		}

		if len(buffer) > 0 {
			inChan <- buffer
			batches++
		}
		close(inChan)
		log.Println("Sent", batches, "batches")
	}()

	go func() {
		wg.Wait()
		close(outChan)
	}()

	acum := 0.0
	size := 0
	for res := range outChan {
		// log.Println("Received", res)
		acum += res
		size++
	}

	fmt.Println("Mean is", acum/float64(size))
}

func processPipelineSequential() {
	fileName := "./data.csv.gz"
	log.Println("Reading from file", fileName)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal("Error decoding gzip", err)
	}
	csvr := csv.NewReader(gr)

	acum := 0
	records := 0
	for {
		row, err := csvr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error reading csv", err)
		}

		col1, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatal("Error converting col1 to int", err)
		}
		col2, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal("Error converting col1 to int", err)
		}

		res := (col1 * col2)
		res = res * res
		acum += res
		records++
	}

	fmt.Println("Mean is", float64(acum)/float64(records))
}

func TestProcessPipeline(t *testing.T) {
	t.Run("Process pipeline", func(t *testing.T) {
		start := time.Now()
		processPipeline()
		fmt.Println("Took", time.Since(start))
	})
}

func TestProcessPipelineSequential(t *testing.T) {
	t.Run("Process pipeline sequential", func(t *testing.T) {
		start := time.Now()
		processPipelineSequential()
		fmt.Println("Took", time.Since(start))
	})
}
