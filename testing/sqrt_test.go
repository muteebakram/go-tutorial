package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
)

// func name for test must start with Test***
func TestSingleValues(t *testing.T) {
	calVal, err := sqrt(4)
	if err != nil {
		t.Fatal("sqrt() error failed.", err)
	}

	actVal := math.Sqrt(4)
	if actVal-calVal > 0.000001 {
		t.Fatal("sqrt() give wrong answer.")
	}

}

func almostEqual(v1, v2 float64) bool {
	return math.Abs(v1-v2) <= 0.001
}

func TestManyValues(t *testing.T) {
	file, err := os.Open("sqrt.csv")
	if err != nil {
		t.Fatal("Cannot open", err)
	}

	rdr := csv.NewReader(file)
	for {
		record, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("error reading cases file - %s", err)
		}
		val, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			t.Fatalf("bad value - %s", record[0])
		}

		expected, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			t.Fatalf("bad value - %s", record[1])
		}

		t.Run(fmt.Sprintf("%f", val), func(t *testing.T) {
			out, err := sqrt(val)
			if err != nil {
				t.Fatal(err)
			}

			if !almostEqual(out, expected) {
				t.Fatalf("%f != %f", out, expected)
			}
		})
	}
}

func TestNegativeErr(t *testing.T) {
	val, err := sqrt(-4)
	if err == nil {
		t.Fatal("Negative test failed: ", err, ", value: ", val)
	}
}

// func name for 'Benchmark' must start with Benchmark***
// Run benchmark without test: go test -v -bench . -run TTT
// CPU Profile:  go test -v -bench . -run TTT -cpuprofile cpu.out
// Read the profile:  go tool pprof cpu.out
// Then list sqrt
func BenchmarkSqrt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		val, err := sqrt(float64(n))
		if err != nil {
			b.Fatal("sqrt() error failed.", err, ", va;ue: ", val)
		}
	}
}
