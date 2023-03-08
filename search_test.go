package gomarkov

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
)

var slice = getSliceFromFile("output.txt")

func getSliceFromFile(filename string) []string {
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return strings.Split(string(f), "\n")
}

func simpleLinearSearch(slice []string, key string) int {
	for i, val := range slice {
		if val == key {
			return i
		}
	}
	return -1
}

func ExampleConcurrentSearchClosest() {
	println(ConcurrentSearchClosest(slice, "8 марта", runtime.NumCPU()))
}

func BenchmarkLinearSearch(b *testing.B) {
	key := "Сталин!"
	for i := 0; i < b.N; i++ {
		simpleLinearSearch(slice, key)
	}
}

func BenchmarkConcurrentLinearSearch(b *testing.B) {
	key := "Сталин!"
	coreNums := []int{2, 4, 8, 16, 32, 64, runtime.NumCPU()}
	for _, coreNum := range coreNums {
		b.Run(fmt.Sprintf("Concurrent-%v", coreNum), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConcurrentSearchClosest(slice, key, 2)
			}
		})
	}
}
