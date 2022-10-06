package main

import (
	"fmt"
	"testing"
)

var testCriteria_sequential = []struct {
	start int
	end   int
}{
	{start: 0, end: 10},
	{start: 0, end: 100},
	{start: 0, end: 1000},
	{start: 0, end: 10000},
	{start: 0, end: 100000},
	{start: 0, end: 1000000},
}

var testCriteria_parallel = []struct {
	start           int
	end             int
	goroutine_count int
}{
	// 10 Goroutines
	{start: 0, end: 10000, goroutine_count: 10},
	{start: 0, end: 1000000, goroutine_count: 10},
	// 100 Goroutines
	{start: 0, end: 10000, goroutine_count: 100},
	{start: 0, end: 1000000, goroutine_count: 100},
	// 1000 Goroutines
	{start: 0, end: 10000, goroutine_count: 1000},
	{start: 0, end: 1000000, goroutine_count: 1000},
}

func BenchmarkSequentialFizzbuzz(b *testing.B) {
	for _, value := range testCriteria_sequential {
		b.Run(fmt.Sprintf("START[%d] END[%d]", value.start, value.end), func(b *testing.B) {
			fizzbuzz_sequential(value.start, value.end)
		})
	}
}

func BenchmarkParallelFizzbuzz(b *testing.B) {
	for _, value := range testCriteria_parallel {
		b.Run(fmt.Sprintf("START[%d] END[%d] GOROUTINES[%d]", value.start, value.end, value.goroutine_count), func(b *testing.B) {
			parallel_fizzbuzz(value.start, value.end, value.goroutine_count)
		})
	}
}

func BenchmarkWorkerPoolFizzbuzz(b *testing.B) {
	for _, value := range testCriteria_parallel {
		b.Run(fmt.Sprintf("START[%d] END[%d] GOROUTINES[%d]", value.start, value.end, value.goroutine_count), func(b *testing.B) {
			worker_pool_fizzbuzz(value.start, value.end, value.goroutine_count)
		})
	}
}
