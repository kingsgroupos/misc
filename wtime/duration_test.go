package wtime

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	var cases = map[string]time.Duration{
		"":       0,
		"1h2m3s": time.Hour + time.Minute*2 + time.Second*3,
		"1d":     time.Hour * 24,
		"2d":     time.Hour * 24 * 2,
		"1d1h":   time.Hour*24 + time.Hour,
		"2d2s":   time.Hour*24*2 + time.Second*2,
		"d":      0,
	}
	for k, v := range cases {
		d, err := ParseDuration(k)
		if d != v {
			t.Fatalf("unexpected result. k: %s, d: %s, v: %s, err: %+s", k, d, v, err)
		}
	}
}
