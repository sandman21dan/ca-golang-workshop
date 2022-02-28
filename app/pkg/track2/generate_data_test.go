package pipeline

import (
	"compress/gzip"
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func generateData() {
	log.Println("Generating data")
	values := 10_000_000

	rand.Seed(42)

	f, err := os.Create("./data.csv.gz")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer f.Close()

	gw := gzip.NewWriter(f)
	defer gw.Close()
	csvw := csv.NewWriter(gw)
	defer csvw.Flush()

	for i := 0; i < values; i++ {
		csvw.Write([]string{
			strconv.Itoa(rand.Intn(1_000_000)),
			strconv.Itoa(rand.Intn(1_000_000)),
		})
	}

	log.Println("Done")
}

func TestGenerateData(t *testing.T) {
	t.Run("Generate random data", func(t *testing.T) {
		generateData()
	})
}
