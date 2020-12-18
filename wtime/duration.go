package wtime

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"gitlab-ee.funplus.io/watcher/watcher/misc/wunsafe"
)

const Day = 24 * time.Hour

func ParseDuration(s string) (time.Duration, error) {
	idx := strings.IndexRune(s, 'd')
	if idx < 0 {
		return time.ParseDuration(s)
	}

	days, err := strconv.Atoi(s[:idx])
	if err != nil {
		return 0, errors.New("time: invalid duration " + s)
	}
	d1 := Day * time.Duration(days)
	if idx+1 >= len(s) {
		return d1, nil
	}
	d2, err := time.ParseDuration(s[idx+1:])
	if err != nil {
		return 0, err
	}

	return d1 + d2, nil
}

type Duration time.Duration

func (d Duration) D() time.Duration {
	return time.Duration(d)
}

func (d Duration) MarshalJSON() ([]byte, error) {
	s := d.String()
	n := len(s) + 2
	b := make([]byte, n)
	copy(b[1:], wunsafe.StringToBytes(s))
	b[0], b[n-1] = '"', '"'
	return b, nil
}

func (this *Duration) UnmarshalJSON(data []byte) error {
	n := len(data)
	if n < 2 {
		return nil
	}
	d, err := ParseDuration(string(data[1 : n-1]))
	if err != nil {
		return err
	}
	*this = Duration(d)
	return nil
}

func (d Duration) Hours() float64 {
	return time.Duration(d).Hours()
}

func (d Duration) Microseconds() int64 {
	return time.Duration(d).Microseconds()
}

func (d Duration) Milliseconds() int64 {
	return time.Duration(d).Milliseconds()
}

func (d Duration) Minutes() float64 {
	return time.Duration(d).Minutes()
}

func (d Duration) Nanoseconds() int64 {
	return time.Duration(d).Nanoseconds()
}

func (d Duration) Round(m time.Duration) time.Duration {
	return time.Duration(d).Round(m)
}

func (d Duration) Seconds() float64 {
	return time.Duration(d).Seconds()
}

func (d Duration) String() string {
	return time.Duration(d).String()
}

func (d Duration) Truncate(m time.Duration) time.Duration {
	return time.Duration(d).Truncate(m)
}
