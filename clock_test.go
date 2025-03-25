package clock_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/michalkurzeja/go-clock/v2"
	"github.com/michalkurzeja/go-clock/v2/mocks"
)

func TestMock(t *testing.T) {
	now := time.Date(2019, time.September, 30, 14, 30, 00, 00, time.UTC)

	t.Run("inner test with fake clock", func(t *testing.T) {
		t.Parallel()

		clk := clock.Mock(t, now)

		time.Sleep(time.Nanosecond) // We just want to be sure that ANY time has passed.
		assert.Equal(t, now, clock.Now())

		clk.Advance(time.Second)
		now = now.Add(time.Second)

		assert.Equal(t, now, clock.Now())
	})

	assert.NotEqual(t, now, clock.Now(), "expected the real time to move on")
}

func TestAfter(t *testing.T) {
	d := time.Second
	ch := make(<-chan time.Time)

	clk := mocks.NewClock(t)
	clk.EXPECT().After(d).Return(ch)
	clock.Set(clk)

	got := clock.After(d)

	assert.Equal(t, ch, got)
}

func TestAfterFunc(t *testing.T) {
	d := time.Second
	f := func() {}

	timer := mocks.NewTimer(t)
	clk := mocks.NewClock(t)
	clk.EXPECT().AfterFunc(d, mock.AnythingOfType("func()")).Return(timer)
	clock.Set(clk)

	got := clock.AfterFunc(d, f)

	assert.Equal(t, timer, got)
}

func TestNow(t *testing.T) {
	now := time.Now()

	clk := mocks.NewClock(t)
	clk.EXPECT().Now().Return(now)
	clock.Set(clk)

	got := clock.Now()

	assert.Equal(t, now, got)
}

func TestSince(t *testing.T) {
	now := time.Now()
	d := time.Second

	clk := mocks.NewClock(t)
	clk.EXPECT().Since(now).Return(d)
	clock.Set(clk)

	got := clock.Since(now)

	assert.Equal(t, d, got)
}

func TestSleep(t *testing.T) {
	d := time.Second

	clk := mocks.NewClock(t)
	clk.EXPECT().Sleep(d).Return()
	clock.Set(clk)

	clock.Sleep(d)
}

func TestTick(t *testing.T) {
	d := time.Second
	ch := make(<-chan time.Time)

	ticker := mocks.NewTicker(t)
	ticker.EXPECT().Chan().Return(ch)
	clk := mocks.NewClock(t)
	clk.EXPECT().NewTicker(d).Return(ticker)
	clock.Set(clk)

	got := clock.Tick(d)

	assert.Equal(t, ch, got)
}

func TestNewTicker(t *testing.T) {
	d := time.Second
	ticker := mocks.NewTicker(t)

	clk := mocks.NewClock(t)
	clk.EXPECT().NewTicker(d).Return(ticker)
	clock.Set(clk)

	got := clock.NewTicker(d)

	assert.Equal(t, ticker, got)
}

func TestTimer(t *testing.T) {
	d := time.Second

	timer := mocks.NewTimer(t)
	clk := mocks.NewClock(t)
	clk.EXPECT().NewTimer(d).Return(timer)
	clock.Set(clk)

	got := clock.NewTimer(d)

	assert.Equal(t, timer, got)
}
