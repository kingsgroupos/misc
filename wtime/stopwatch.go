package wtime

import (
	"fmt"
	"time"
)

type lapRecord struct {
	name string
	d    time.Duration
}

// Stopwatch helps you measure elapsed time.
type Stopwatch struct {
	startTime time.Time
	laps      []lapRecord
}

// NewStopwatch creates a new stopwatch instance and starts measuring elapsed time.
func NewStopwatch() Stopwatch {
	return Stopwatch{
		startTime: time.Now(),
	}
}

// Elapsed returns the total elapsed time.
func (sw Stopwatch) Elapsed() time.Duration {
	return time.Since(sw.startTime)
}

// ElapsedSeconds returns the total elapsed seconds.
func (sw Stopwatch) ElapsedSeconds() float64 {
	duration := time.Since(sw.startTime)
	return duration.Seconds()
}

// ElapsedMilliseconds returns the total elapsed milliseconds.
func (sw Stopwatch) ElapsedMilliseconds() float64 {
	duration := time.Since(sw.startTime)
	return duration.Seconds() * 1000
}

func (this *Stopwatch) Lap(name string) {
	this.laps = append(this.laps, lapRecord{
		name: name,
		d:    this.Elapsed(),
	})
}

func (sw Stopwatch) ElapsedSince(lapName string) time.Duration {
	d := sw.Elapsed()
	for i := len(sw.laps) - 1; i >= 0; i-- {
		if sw.laps[i].name == lapName {
			return d - sw.laps[i].d
		}
	}

	panic(fmt.Errorf("cannot find the lap %s in the records", lapName))
}

func (sw Stopwatch) ElapsedSecondsSince(lapName string) float64 {
	d := sw.ElapsedSince(lapName)
	return d.Seconds()
}

func (sw Stopwatch) ElapsedMillisecondsSince(lapName string) float64 {
	d := sw.ElapsedSince(lapName)
	return d.Seconds() * 1000
}

func (sw Stopwatch) RangeReversedLaps(f func(name string, d time.Duration) bool) {
	for i := len(sw.laps) - 1; i >= 0; i-- {
		lap := sw.laps[i]
		if !f(lap.name, lap.d) {
			return
		}
	}
}
