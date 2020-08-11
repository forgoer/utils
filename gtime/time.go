package gtime

import (
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(Layout))
	return []byte(formatted), nil
}
