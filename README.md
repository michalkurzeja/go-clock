# go-clock
A small package that offers testable time functions.

## Why?
It is hard to properly test the code that performs time operations relative to the current time
or that depends on the time flow, e.g. with `time.Now()` or `time.After(d)` calls.

`go-clock` provides a package-level clock object that can easily be swapped for a configurable mock in your tests.
The package also offers some commonly-used functions from the `time` package that use the clock.

## Installation
```shell script
go get github.com/michalkurzeja/go-clock/v2
```

## Usage
In your code, simply use the `go-clock` functions for time retrieval instead of the standard `time` package:

```go
package "main"

import (
	"time"
	
	"github.com/michalkurzeja/go-clock/v2"
)

func main() {
	now := clock.Now()            // Instead of `time.Now()`
	since := clock.Since(now)     // Instead of `time.Since()`
	c := clock.After(time.Second) // Instead of `time.After(time.Second)`
}
```

In your tests, you can mock the clock to get predictable time output:

```go
package main_test

import (
	"testing"
	"time"
	
	"github.com/michalkurzeja/go-clock/v2"
)

func TestSomething(t *testing.T) {
	fakeNow := time.Now()
	
	clk := clock.Mock(t, fakeNow) // `clock.Now()` will always return `fakeNow` time.
	// After the test ends, the clock will automatically be restored to real clock.
	
	clk.Advance(time.Second) // Advances the fake clock's time by a second.
}
```

`go-clock` uses the clock implementation from the [jonboulle/clockwork](https://github.com/jonboulle/clockwork) package.
For more details on how it works, please see it's docs.
