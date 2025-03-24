## rpsmeter

A lightweight, high-performance library for measuring request rates per second in Go applications.

### Features

- Efficient tracking of request rates using atomic operations
- Low memory footprint with fixed-size bucket storage
- Thread-safe implementation
- Ability to retrieve per-second statistics for the last 10 seconds
- Results can be fed infrequently to [hdrhistogram-go](https://github.com/HdrHistogram/hdrhistogram-go) for further analysis

![Blank diagram](https://github.com/user-attachments/assets/45f15804-868b-40aa-ba51-dcb42d329336)

### Installation

```bash
go get github.com/makasim/rpsmeter
```

### Usage

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/makasim/rpsmeter"
)

func main() {
	rpsmeter.Init()
	
    // Create a new meter
    var meter rpsmeter.Meter
    
    // Record requests as they happen
    meter.Record()
    
    // Get the total and the breakdown count of requests for the last 10 seconds
    total, breakdown := meter.Result()
    fmt.Println(counts)
}
```

### Example

See the [examples/main.go](examples/main.go) for a complete example of how to use the meter.

### Benchmarks

```bash
goos: darwin
goarch: arm64
pkg: github.com/makasim/rpsmeter
cpu: Apple M1 Pro
BenchmarkMeter_Record
BenchmarkMeter_Record-10             	16777956	        62.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkMeter_Parallel_Record
BenchmarkMeter_Parallel_Record-10    	 7141261	       160.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkMeter_Result
BenchmarkMeter_Result-10             	18844047	        63.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkMeter_Parallel_Result
BenchmarkMeter_Parallel_Result-10    	13420922	        91.05 ns/op	       0 B/op	       0 allocs/op
PASS
```

### License

MIT
