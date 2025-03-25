package clock

import (
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
)

type Clock clockwork.Clock

type Timer clockwork.Timer

type Ticker clockwork.Ticker

// clk represents a global clock.
var clk Clock = clockwork.NewRealClock()

// Mock replaces the clk with a mock frozen at the given time and returns it.
func Mock(t *testing.T, now time.Time) *clockwork.FakeClock {
	fake := clockwork.NewFakeClockAt(now)
	Set(fake)
	t.Cleanup(func() {
		clk = clockwork.NewRealClock()
	})
	return fake
}

// Set sets the global clock.
func Set(c Clock) {
	clk = c
}

// After waits for the duration to elapse and then sends the current time
func After(d time.Duration) <-chan time.Time {
	return clk.After(d)
}

// AfterFunc waits for the duration to elapse and then calls f in its own goroutine.
func AfterFunc(d time.Duration, f func()) Timer {
	return clk.AfterFunc(d, f)
}

// Now returns the current local time.
func Now() time.Time {
	return clk.Now()
}

// Since returns the time elapsed since t.
func Since(t time.Time) time.Duration {
	return clk.Since(t)
}

// Sleep pauses the current goroutine for at least the duration d.
func Sleep(d time.Duration) {
	clk.Sleep(d)
}

// Tick is a convenience wrapper for NewTicker providing access to the ticking channel only.
func Tick(d time.Duration) <-chan time.Time {
	return clk.NewTicker(d).Chan()
}

// NewTicker returns a new Ticker containing a channel that will send the
// time with a period specified by the duration argument.
func NewTicker(d time.Duration) clockwork.Ticker {
	return clk.NewTicker(d)
}

// NewTimer creates a new Timer that will send the current time on its channel after at least duration d.
func NewTimer(d time.Duration) clockwork.Timer {
	return clk.NewTimer(d)
}
