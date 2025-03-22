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

### License

MIT