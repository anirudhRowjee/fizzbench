# FizzBench

Benchmarking Sequential vs Parallelized implementations of the all-time classic Fizzbuzz.

Some example results -
```
λ go run main.go
2022/10/04 13:31:09 FIZZBUZZ BENCHMARK REPORT
2022/10/04 13:31:09 GOROUTINES: 10000 | START: 0 | END: 1000000
2022/10/04 13:31:09 SEQUENTIAL Execution took (AVG over 10 runs) 0.212155 sec
2022/10/04 13:31:09 PARALLEL Execution took (AVG over 10 runs) 0.055120 sec
2022/10/04 13:31:09 Average Speedup -> 3.848961x
```

```
λ go run main.go
2022/10/04 13:17:23 FIZZBUZZ BENCHMARK REPORT
2022/10/04 13:17:23 GOROUTINES: 10000 | START: 0 | END: 1000000
2022/10/04 13:17:23 SEQUENTIAL Execution took (AVG over 10 runs) 0.216635 sec
2022/10/04 13:17:23 PARALLEL Execution took (AVG over 10 runs) 0.057681 sec
2022/10/04 13:17:23 Average Speedup -> 3.755725x
```