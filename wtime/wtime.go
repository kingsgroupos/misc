package wtime

import "time"

const (
	DateTimeFormat                   = "2006-01-02 15:04:05"
	CompactDateTimeFormat            = "20060102150405"
	CompactDateTimeFormatWithoutYear = "0102150405"
	DateFormat                       = "2006-01-02"
	TimeFormat                       = "15:04:05"
)

var (
	Unix1970 = time.Unix(0, 0).UTC()
)

func UnixMilli(t time.Time) int64 {
	return t.Unix()*1e3 + int64(t.Nanosecond()/1e6)
}

func UnixMicro(t time.Time) int64 {
	return t.Unix()*1e6 + int64(t.Nanosecond()/1e3)
}

func CompareDateUTC(t1, t2 time.Time) int {
	v1 := t1.Unix() / 86400
	v2 := t2.Unix() / 86400
	switch {
	case v1 == v2:
		return 0
	case v1 < v2:
		return -1
	default:
		return 1
	}
}

func TimestampPass(timestamp, t int64) bool {
	return t <= timestamp
}
