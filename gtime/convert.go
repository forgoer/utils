package gtime

import (
	"strings"
	"time"
)

func Format(t Time) string {
	return t.Format(Layout)
}

func Parse(s string) (time.Time, error) {
	return time.ParseInLocation(Layout, s, time.Local)
}

func Sequence() string {
	now := time.Now().Format("2006.01.02.15.04.05.000")
	timeSlot := strings.Replace(now, ".", "", -1)
	return timeSlot
}
