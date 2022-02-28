package pipeline

import (
	"fmt"
	"testing"
	"time"
)

func processPipeline() {
	fileName := "./data.csv.gz"
}

func TestProcessPipeline(t *testing.T) {
	t.Run("Process pipeline", func(t *testing.T) {
		start := time.Now()
		processPipeline()
		fmt.Println("Took", time.Since(start))
	})
}
