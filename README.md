# go-clock
[![Build Status](https://travis-ci.com/msales/go-clock.svg?token=jnuRixQ5JT2Tqqcethpp&branch=master)](https://travis-ci.com/msales/go-clock)
[![Coverage Status](https://coveralls.io/repos/github/msales/go-clock/badge.svg?branch=master&t=x6DuOO)](https://coveralls.io/github/msales/go-clock?branch=master)

A small package that offers testable time functions.

## Why?
It is hard to properly test the code that performs time operations relative to the current time
or that depends on the time flow, e.g. with `time.Now()` or `time.After(d)` calls.

`go-clock` provides a package-level Clock object that can easily be swapped for a configurable mock in your tests.
The package also offers some commonly-used functions from the `time` package that use the `Clock`.

## Installation
```shell script
go get github.com/michalkurzeja/go-clock
```

## Usage
In your code, simply use the `go-clock` functions for time retrieval instead of the standard `time` package:

```go
import "github.com/michalkurzeja/go-clock"

now := clock.Now() // Instead of `time.Now()`
since := clock.Since(now) // Instead of `time.Since()`
c := clock.After(time.Second) // Instead of `time.After(time.Second)`
```

In your tests, you can mock the clock to get predictable time output:

```go
fakeNow := time.Date(...)

mock := clock.Mock(fakeNow) // `clock.Now()` will always return `fakeNow` time.
defer clock.Restore()

mock.Add(time.Second) // Advances the fake clock's time by a second.
```

`go-clock` uses the clock implementation from the [benbjohnson/clock](https://github.com/benbjohnson/clock) package.
For more details on the usage of the clock, please see it's docs.
