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
    // Create a new meter
    var meter rpsmeter.Meter
    
    // Record requests as they happen
    meter.Record()
    
    // Get the count of requests for the last 10 seconds
    counts := meter.LastTenSeconds()
    fmt.Println(counts)
}
```

### Example

See the [examples/main.go](examples/main.go) for a complete example of how to use the meter.

### Benchmarks

```bash
$go test -run=none -bench=. --memprofile=out.mem 
2025/03/22 18:36:39 489
goos: darwin
goarch: arm64
pkg: github.com/makasim/rpsmeter
cpu: Apple M1 Pro
BenchmarkMeter_Record-10             	170864402	         6.965 ns/op	       0 B/op	       0 allocs/op
BenchmarkMeter_Parallel_Record-10    	 7997917	       145.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMeter_Result-10             	67978878	        17.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkMeter_Parallel_Result-10    	510158802	         2.170 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/makasim/rpsmeter	5.302s
```

### License

MIT