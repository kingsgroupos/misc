// BSD 3-Clause License
//
// Copyright (c) 2020, Kingsgroup
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

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
