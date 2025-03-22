## rpsmeter

A lightweight, high-performance library for measuring request rates per second in Go applications.

### Features

- Efficient tracking of request rates using atomic operations
- Low memory footprint with fixed-size bucket storage
- Thread-safe implementation
- Ability to retrieve per-second statistics for the last 10 seconds

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
BenchmarkMeter_Record
BenchmarkMeter_Record-10             	102875260	        11.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkMeter_Parallel_Record
BenchmarkMeter_Parallel_Record-10    	 5847054	       244.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkMeter_Result
BenchmarkMeter_Result-10             	52411759	        22.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkMeter_Parallel_Result
BenchmarkMeter_Parallel_Result-10    	440293185	         2.799 ns/op	       0 B/op	       0 allocs/op
```

### License

MIT