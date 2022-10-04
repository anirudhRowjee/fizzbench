package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

func fizzbuzz_sequential(start, end int) []string {

	// declare the array we're using to keep the strings in
	var fb []string

	for i := start; i <= end; i++ {

		// keep a variable to check if the number is divisible by 3 or 5
		divisible := 0

		// no else used here to ensure numbers divisbile by 3 and 5 show up as "fizzbuzz"
		if i%3 == 0 {
			fb = append(fb, "fizz")
			divisible = 1
		}
		if i%5 == 0 {
			fb = append(fb, "buzz")
			divisible = 1
		}

		// if we haven't appended anything, push the number itself
		if divisible == 0 {
			fb = append(fb, strconv.Itoa(i))
		}

		fb = append(fb, "\n")

	}
	return fb
}

func process_result(c <-chan []string) [][]string {
	// function to process the results of the channel
	var fb [][]string
	for segment := range c {
		fb = append(fb, segment)
	}
	// fmt.Println("Length of fb ->", len(fb))
	return fb
}

func parallel_fizzbuzz(start, end, degree int) [][]string {
	// @param Degree - number of goroutines to spawn

	// declare the array we're using to keep the strings in
	// var fb [][]string

	// segment the start and end into ranges based on the degree
	// @variable jump_div_check is the increment we use to move between start and end ranges
	// For now, we ensure that jump_div_check is an whole number.
	jump_div_check := (end - start) % degree

	if jump_div_check != 0 {
		log.Fatalln("The degree must perfectly divide the number of entries")
	}

	jump := (end - start) / degree
	// fmt.Println("Jump -> ", jump)

	// declare the iteration variable
	iter_start := start

	// create the channel we'll be communicating on
	number_channel := make(chan []string)
	// create the waitgroup
	var wg sync.WaitGroup

	// simple while-loop
	for {
		if iter_start < end {

			range_start := iter_start
			range_end := iter_start + jump

			// log.Printf("Launching Goroutine for Fizzbuzz (%d, %d)\n", range_start+1, range_end)

			// launch the goroutine
			wg.Add(1)
			go func(start, end int) {
				number_channel <- fizzbuzz_sequential(start+1, end)
				wg.Done()
			}(range_start, range_end)

			iter_start += jump

		} else {
			break
		}
	}

	go func() {
		wg.Wait()
		close(number_channel)
	}()

	return process_result(number_channel)
}

func main() {

	// 1 to @param limit for fizzbuzz
	// limit, err := strconv.Atoi(os.Args[1])
	// // 1 - write to stdout, 2 - write to string
	// mode, _ := strconv.Atoi(os.Args[2])

	start_number := 0
	limit := 1000000

	runs := 10
	goroutines := 100

	// mode := 3

	// fmt.Println("Hello, world!")
	// fmt.Println("Count to follow till ->", mode)
	seq_time := 0.0
	par_time := 0.0

	for i := 0; i < runs; i++ {

		start := time.Now()
		fizzbuzz_sequential(start_number, limit)
		elapsed := time.Since(start).Seconds()
		seq_time += elapsed

		start_parallel := time.Now()
		parallel_fizzbuzz(start_number, limit, goroutines)
		elapsed_parallel := time.Since(start_parallel).Seconds()
		par_time += elapsed_parallel

	}

	seq_time = seq_time / float64(runs)
	par_time = par_time / float64(runs)

	log.Printf("FIZZBUZZ BENCHMARK REPORT")
	log.Printf("GOROUTINES: %d | START: %d | END: %d", goroutines, start_number, limit)
	log.Printf("SEQUENTIAL Execution took (AVG over %d runs) %f sec", runs, seq_time)
	log.Printf("PARALLEL Execution took (AVG over %d runs) %f sec", runs, par_time)
	log.Printf("Average Speedup -> %fx", seq_time/par_time)
	// fmt.Println(outbuf)

}
