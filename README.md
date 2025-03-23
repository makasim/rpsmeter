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





![Uploading<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:lucid="lucid" width="1242" height="201.38"><g transform="translate(-79 -199.62360346032693)" lucid:page-tab-id="0_0"><path d="M0 0h1500v500H0z" fill="#fff"/><path d="M980 285a6 6 0 0 1 6-6h328a6 6 0 0 1 6 6v109a6 6 0 0 1-6 6H986a6 6 0 0 1-6-6z" fill="#fff"/><path d="M981 285c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm6-6c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm6 6c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.57.45-1 1-1s1 .43 1 1zm0 7.77c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.56.45-1 1-1s1 .44 1 1zm0 7.77c0 .56-.45 1-1 1s-1-.44-1-1c0-.55.45-1 1-1s1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.56.45-1 1-1s1 .44 1 1zm0 7.77c0 .57-.45 1-1 1s-1-.43-1-1c0-.54.45-1 1-1s1 .46 1 1zm0 7.8c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.57.45-1 1-1s1 .43 1 1zm0 7.77c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.56.45-1 1-1s1 .44 1 1zm0 7.77c0 .56-.45 1-1 1s-1-.44-1-1c0-.55.45-1 1-1s1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.56.45-1 1-1s1 .44 1 1zm0 7.77c0 .57-.45 1-1 1s-1-.43-1-1c0-.54.45-1 1-1s1 .46 1 1zm0 7.8c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-6 6c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-6-6c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0-7.8c0 .57-.45 1-1 1s-1-.43-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0-7.8c0 .56-.45 1-1 1s-1-.44-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .55-.45 1-1 1s-1-.45-1-1c0-.56.45-1 1-1s1 .44 1 1zm0-7.8c0 .56-.45 1-1 1s-1-.44-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .54-.45 1-1 1s-1-.46-1-1c0-.57.45-1 1-1s1 .43 1 1zm0-7.8c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0-7.8c0 .57-.45 1-1 1s-1-.43-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0-7.8c0 .56-.45 1-1 1s-1-.44-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .55-.45 1-1 1s-1-.45-1-1c0-.56.45-1 1-1s1 .44 1 1zm0-7.8c0 .56-.45 1-1 1s-1-.44-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .54-.45 1-1 1s-1-.46-1-1c0-.57.45-1 1-1s1 .43 1 1z" fill="#3a414a"/><use xlink:href="#a" transform="matrix(1,0,0,1,991.9999999999999,291) translate(115.68771701388889 57.90277777777778)"/><path d="M80 285a6 6 0 0 1 6-6h488a6 6 0 0 1 6 6v109a6 6 0 0 1-6 6H86a6 6 0 0 1-6-6z" fill="#fff"/><path d="M81 285c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm6-6c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm6 6c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.57.45-1 1-1s1 .43 1 1zm0 7.77c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.56.45-1 1-1s1 .44 1 1zm0 7.77c0 .56-.45 1-1 1s-1-.44-1-1c0-.55.45-1 1-1s1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.56.45-1 1-1s1 .44 1 1zm0 7.77c0 .57-.45 1-1 1s-1-.43-1-1c0-.54.45-1 1-1s1 .46 1 1zm0 7.8c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.57.45-1 1-1s1 .43 1 1zm0 7.77c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.56.45-1 1-1s1 .44 1 1zm0 7.77c0 .56-.45 1-1 1s-1-.44-1-1c0-.55.45-1 1-1s1 .45 1 1zm0 7.8c0 .54-.45 1-1 1s-1-.46-1-1c0-.56.45-1 1-1s1 .44 1 1zm0 7.77c0 .57-.45 1-1 1s-1-.43-1-1c0-.54.45-1 1-1s1 .46 1 1zm0 7.8c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-6 6c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-8 0c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm-6-6c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0-7.8c0 .57-.45 1-1 1s-1-.43-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0-7.8c0 .56-.45 1-1 1s-1-.44-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .55-.45 1-1 1s-1-.45-1-1c0-.56.45-1 1-1s1 .44 1 1zm0-7.8c0 .56-.45 1-1 1s-1-.44-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .54-.45 1-1 1s-1-.46-1-1c0-.57.45-1 1-1s1 .43 1 1zm0-7.8c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0-7.8c0 .57-.45 1-1 1s-1-.43-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .55-.45 1-1 1s-1-.45-1-1 .45-1 1-1 1 .45 1 1zm0-7.8c0 .56-.45 1-1 1s-1-.44-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .55-.45 1-1 1s-1-.45-1-1c0-.56.45-1 1-1s1 .44 1 1zm0-7.8c0 .56-.45 1-1 1s-1-.44-1-1c0-.54.45-1 1-1s1 .46 1 1zm0-7.77c0 .54-.45 1-1 1s-1-.46-1-1c0-.57.45-1 1-1s1 .43 1 1z" fill="#3a414a"/><use xlink:href="#a" transform="matrix(1,0,0,1,92,291) translate(195.68771701388889 57.90277777777778)"/><path d="M100 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#b" transform="matrix(1,0,0,1,112,312) translate(21.245442708333336 35.52777777777778)"/><path d="M180 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#c" transform="matrix(1,0,0,1,192,312) translate(21.223741319444443 35.52777777777778)"/><path d="M260 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#d" transform="matrix(1,0,0,1,272,312) translate(21.712022569444443 35.52777777777778)"/><path d="M340 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#e" transform="matrix(1,0,0,1,352,312) translate(20.99045138888889 35.52777777777778)"/><path d="M420 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#f" transform="matrix(1,0,0,1,432,312) translate(21.126085069444443 35.52777777777778)"/><path d="M500 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#g" transform="matrix(1,0,0,1,512,312) translate(21.13693576388889 35.52777777777778)"/><path d="M580 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#c" transform="matrix(1,0,0,1,592,312) translate(21.223741319444443 35.52777777777778)"/><path d="M660 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#b" transform="matrix(1,0,0,1,672,312) translate(21.245442708333336 35.52777777777778)"/><path d="M740 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#e" transform="matrix(1,0,0,1,752,312) translate(20.99045138888889 35.52777777777778)"/><path d="M820 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#e" transform="matrix(1,0,0,1,832,312) translate(20.99045138888889 35.52777777777778)"/><path d="M900 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#d" transform="matrix(1,0,0,1,912,312) translate(21.712022569444443 35.52777777777778)"/><path d="M980 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#c" transform="matrix(1,0,0,1,992,312) translate(21.223741319444443 35.52777777777778)"/><path d="M1060 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#g" transform="matrix(1,0,0,1,1072,312) translate(21.13693576388889 35.52777777777778)"/><path d="M1140 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#h" transform="matrix(1,0,0,1,1152,312) translate(16.704427083333336 35.52777777777778)"/><path d="M1220 306a6 6 0 0 1 6-6h68a6 6 0 0 1 6 6v68a6 6 0 0 1-6 6h-68a6 6 0 0 1-6-6z" stroke="#3a414a" stroke-width="2" fill="#fff"/><use xlink:href="#f" transform="matrix(1,0,0,1,1232,312) translate(21.126085069444443 35.52777777777778)"/><path d="M700 221v59.5M700 221.03V220" stroke="#3a414a" stroke-width="2" fill="none"/><path d="M700 295.76l-4.63-14.26h9.26z" stroke="#3a414a" stroke-width="2" fill="#3a414a"/><use xlink:href="#i" transform="matrix(1,0,0,1,708,199.62360346032693) translate(0 17.22222222222222)"/><use xlink:href="#j" transform="matrix(1,0,0,1,708,199.62360346032693) translate(0 38.732638888888886)"/><path d="M860 221v59.5M860 221.03V220" stroke="#3a414a" stroke-width="2" fill="none"/><path d="M860 295.76l-4.63-14.26h9.26z" stroke="#3a414a" stroke-width="2" fill="#3a414a"/><use xlink:href="#k" transform="matrix(1,0,0,1,868,209.24479166666666) translate(0 17.22222222222222)"/><path d="M578.56 260H476.38M578.55 260h.5" stroke="#3a414a" fill="none"/><path d="M461.62 260l14.26-4.63v9.26z" stroke="#3a414a" fill="#3a414a"/><use xlink:href="#l" transform="matrix(1,0,0,1,480.796607001497,234.48958333333334) translate(0 17.22222222222222)"/><use xlink:href="#m" transform="matrix(1,0,0,1,480.796607001497,234.48958333333334) translate(8.42013888888889 17.22222222222222)"/><use xlink:href="#n" transform="matrix(1,0,0,1,480.796607001497,234.48958333333334) translate(49.470486111111114 17.22222222222222)"/><path d="M1319.5 260h-102.17M1319.5 260h.5" stroke="#3a414a" fill="none"/><path d="M1202.56 260l14.27-4.63v9.26z" stroke="#3a414a" fill="#3a414a"/><defs><path d="M180 0v-1490h510c348 0 508 209 508 474 0 266-160 477-507 477H370V0H180zm190-706h312c236 0 327-133 327-310 0-176-91-307-329-307H370v617" id="o"/><path d="M158 0v-1118h174v172h12c41-113 157-188 290-188 26 0 70 2 91 3v181c-11-2-60-10-108-10-161 0-279 109-279 260V0H158" id="p"/><path d="M613 24c-304 0-509-231-509-576 0-350 205-580 509-580 305 0 511 230 511 580 0 345-206 576-511 576zm0-161c226 0 329-195 329-415 0-222-103-419-329-419-223 0-326 196-326 419 0 220 103 415 326 415" id="q"/><path d="M613 24c-304 0-509-231-509-576 0-350 205-580 509-580 216 0 392 114 453 309l-173 49c-33-115-133-197-280-197-223 0-326 196-326 419 0 220 103 415 326 415 150 0 252-85 285-206l172 49C1010-95 832 24 613 24" id="r"/><path d="M628 24c-324 0-524-230-524-574 0-343 198-582 503-582 237 0 487 146 487 559v75H286c9 234 145 362 343 362 132 0 231-58 273-172l174 48C1024-91 857 24 628 24zM287-650h624c-17-190-120-322-304-322-192 0-309 151-320 322" id="s"/><path d="M538 24C308 24 148-78 108-271l171-41c32 123 123 178 257 178 156 0 256-77 256-169 0-77-54-128-164-154l-186-44c-203-48-300-148-300-305 0-192 176-326 414-326 230 0 351 112 402 269l-163 42c-31-80-94-158-238-158-133 0-233 69-233 162 0 83 57 129 188 160l169 40c203 48 298 149 298 302 0 196-179 339-441 339" id="t"/><g id="a"><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,0,0)" xlink:href="#o"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,14.192708333333332,0)" xlink:href="#p"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,22.189670138888886,0)" xlink:href="#q"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,35.514322916666664,0)" xlink:href="#r"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,48.209635416666664,0)" xlink:href="#s"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,61.16536458333333,0)" xlink:href="#t"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,72.89496527777777,0)" xlink:href="#t"/></g><path d="M626 20c-262 0-458-168-468-396h184c12 133 134 229 284 229 180 0 311-137 311-326 0-192-136-335-323-335-92 0-196 33-255 78l-178-22 88-738h784v167H429l-51 435h8c61-51 160-87 263-87 273 0 474 211 474 499 0 286-210 496-497 496" id="u"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,0,0)" xlink:href="#u" id="b"/><path d="M154 0v-137l495-537c165-179 249-281 249-418 0-156-121-253-280-253-170 0-278 110-278 278H158c0-264 200-443 465-443 266 0 455 183 455 416 0 161-73 288-336 568L416-179v12h687V0H154" id="v"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,0,0)" xlink:href="#v" id="c"/><path d="M200 0l662-1311v-12H98v-167h963v177L400 0H200" id="w"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,0,0)" xlink:href="#w" id="d"/><path d="M646 20c-332 0-524-278-524-764 0-483 194-766 524-766s524 283 524 766c0 485-191 764-524 764zm0-166c218 0 341-220 341-598 0-380-123-601-341-601s-341 222-341 601c0 378 123 598 341 598" id="x"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,0,0)" xlink:href="#x" id="e"/><path d="M633 20c-303 0-511-173-511-416 0-188 124-348 291-378v-8c-145-37-237-174-237-332 0-227 192-396 457-396 261 0 456 169 456 396 0 158-94 295-235 332v8c162 30 291 190 291 378 0 243-212 416-512 416zm0-165c197 0 322-103 322-261 0-165-138-283-322-283-188 0-324 118-324 283 0 158 123 261 324 261zm0-703c157 0 272-101 272-252 0-149-110-246-272-246-165 0-273 97-273 246 0 151 112 252 273 252" id="y"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,0,0)" xlink:href="#y" id="f"/><path d="M635 20c-292 0-500-160-510-396h192c11 142 145 229 315 229 187 0 323-105 323-260 0-161-125-274-346-274H488v-165h121c174 0 294-100 294-254 0-148-104-245-266-245-152 0-291 85-297 230H157c8-234 222-395 484-395 278 0 448 188 448 400 0 168-95 291-247 336v12c190 31 301 169 301 357 0 244-216 425-508 425" id="z"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,0,0)" xlink:href="#z" id="g"/><path d="M653-1490V0H466v-1314h-10L96-1047v-204l324-239h233" id="A"/><g id="h"><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,0,0)" xlink:href="#A"/><use transform="matrix(0.010850694444444444,0,0,0.010850694444444444,9.038628472222221,0)" xlink:href="#v"/></g><path fill="#333" d="M135 0v-1490h340l471 754c59 97 122 210 187 365-32-387-14-721-19-1119h312V0h-341C855-382 627-689 427-1113c32 420 16 695 20 1113H135" id="B"/><path fill="#333" d="M628 22C291 22 81-210 81-554c0-346 210-578 547-578 336 0 547 232 547 578 0 344-211 576-547 576zm0-236c160 0 242-146 242-341 0-197-82-341-242-341S387-753 387-555c0 195 81 341 241 341" id="C"/><path fill="#333" d="M361 0L31-1118h315c68 295 127 523 188 845 63-317 128-551 199-845h276c71 296 131 525 196 846 57-321 119-552 186-846h319L1378 0h-305c-73-261-148-486-204-771C813-484 740-261 666 0H361" id="D"/><g id="i"><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,0,0)" xlink:href="#B"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,13.550347222222223,0)" xlink:href="#C"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,24.210069444444446,0)" xlink:href="#D"/></g><path fill="#333" d="M440-1490V0H135v-1490h305" id="E"/><path fill="#333" d="M428-647V0H128v-1118h283l4 231c63-156 176-245 348-245 232 0 384 159 384 421V0H847v-659c0-139-76-222-202-222-128 0-217 86-217 234" id="F"/><path fill="#333" d="M628 22C291 22 81-210 81-554c0-346 210-578 547-578 265 0 455 142 497 372l-279 52c-25-116-99-188-215-188-160 0-244 139-244 341 0 200 84 341 244 341 116 0 193-74 217-195l279 51C1085-123 896 22 628 22" id="G"/><g id="j"><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,0,0)" xlink:href="#E"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,4.991319444444445,0)" xlink:href="#F"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,16.05902777777778,0)" xlink:href="#G"/></g><path fill="#333" d="M135 0v-1490h581c339 0 533 192 533 485 0 204-94 351-264 424L1302 0H964L680-532H440V0H135zm305-779h220c186 0 274-77 274-226 0-150-88-233-275-233H440v459" id="H"/><path fill="#333" d="M633 22C291 22 81-200 81-553c0-345 210-579 537-579 292 0 525 185 525 567v84H378c7 183 109 277 260 277 105 0 184-46 216-132l272 51C1071-99 893 22 633 22zM380-669h474c-15-146-94-237-233-237-143 0-229 99-241 237" id="I"/><path fill="#333" d="M572 22C302 22 117-98 78-306l279-48c29 105 101 156 224 156 114 0 185-48 185-115 0-118-228-130-337-155-215-49-322-150-322-317 0-215 187-347 470-347 265 0 424 116 466 298l-266 47c-24-76-87-130-196-130-98 0-174 46-174 113 0 56 38 95 144 117l201 40c216 44 320 143 320 302 0 220-207 367-500 367" id="J"/><path fill="#333" d="M683-1118v229H474v562c0 71 29 104 102 104 23 0 73-7 95-13l43 225C647 9 580 16 520 16 296 16 174-96 174-301v-588H20v-229h154v-266h300v266h209" id="K"/><g id="k"><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,0,0)" xlink:href="#H"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,11.458333333333334,0)" xlink:href="#I"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,22.048611111111114,0)" xlink:href="#J"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,32.005208333333336,0)" xlink:href="#I"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,42.595486111111114,0)" xlink:href="#K"/></g><path fill="#333" d="M135 0v-1490h305v1237h643V0H135" id="L"/><path fill="#333" d="M440 22C227 22 68-92 68-313c0-250 202-322 422-343 195-21 271-24 271-105 0-95-63-149-175-149-116 0-187 57-211 129l-275-46c57-194 237-305 487-305 240 0 474 108 474 379V0H777v-155h-10C713-51 605 22 440 22zm85-211c140 0 238-93 238-211v-121c-37 26-156 42-222 51-111 16-185 59-185 149 0 87 70 132 169 132" id="M"/><g id="m"><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,0,0)" xlink:href="#L"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,10.052083333333334,0)" xlink:href="#M"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,20.37326388888889,0)" xlink:href="#J"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,30.32986111111111,0)" xlink:href="#K"/></g><g id="n"><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,0,0)" xlink:href="#K"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,6.336805555555556,0)" xlink:href="#I"/><use transform="matrix(0.008680555555555556,0,0,0.008680555555555556,16.927083333333336,0)" xlink:href="#F"/></g></defs></g></svg> Blank diagram.svg…]()
