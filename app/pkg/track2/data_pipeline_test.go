package pipeline

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func processPipeline() {
	fileName := "./data.csv.gz"
	log.Println("Reading from file", fileName)
}

func TestProcessPipeline(t *testing.T) {
	t.Run("Process pipeline", func(t *testing.T) {
		start := time.Now()
		processPipeline()
		fmt.Println("Took", time.Since(start))
	})
}
