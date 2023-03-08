package gomarkov

import (
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"sync"
)

func searchClosest(slice []string, key string) (float64, int) {
	maxI := 0
	maxSimilarity := float64(0)
	for i, val := range slice {
		similarity := strutil.Similarity(val, key, metrics.NewJaroWinkler())
		if similarity > maxSimilarity {
			maxSimilarity = similarity
			maxI = i
		}
	}
	return maxSimilarity, maxI
}

func ConcurrentSearchClosest(slice []string, key string, chunkCount int) int {
	var wg sync.WaitGroup
	chunkSize := len(slice) / chunkCount // split slice into 4 chunks
	chunks := make([][]string, chunkCount)
	for i := 0; i < chunkCount; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == chunkCount-1 {
			end = len(slice)
		}
		chunks[i] = slice[start:end]
	}

	results := make([]struct {
		similarity float64
		i          int
	}, chunkCount)
	wg.Add(chunkCount)
	for i := 0; i < chunkCount; i++ {
		go func(chunk []string, i int) {
			defer wg.Done()
			rS, rI := searchClosest(chunk, key)
			results[i].similarity = rS
			results[i].i = rI + i*chunkSize
		}(chunks[i], i)
	}
	wg.Wait()

	maxI := -1
	maxSimilarity := float64(0)

	for _, s := range results {
		if s.similarity > maxSimilarity {
			maxSimilarity = s.similarity
			maxI = s.i
		}
	}

	return maxI
}
