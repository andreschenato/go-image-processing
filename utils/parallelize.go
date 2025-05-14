package utils

import (
	"log/slog"
	"runtime"
	"sync"
	"time"
)

func parallelizeProcessing(width, height int, process func(startX, endX, height int)) {
	start := time.Now()
	numCPU := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(numCPU)

	chunkSize := max(width/numCPU, 1)

	for i := range numCPU {
		startX := i * chunkSize
		endX := (i + 1) * chunkSize
		if i == numCPU-1 {
			endX = width
		}

		go func(startX, endX int) {
			defer wg.Done()
			process(startX, endX, height)
		}(startX, endX)
	}
	wg.Wait()
	slog.Info(time.Since(start).String())
}