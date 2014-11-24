package xparam

import (
	"fmt"
	"time"
)

//------------------------------------------------------------
// Xparam access for time
//------------------------------------------------------------

// Gets parameter as time. Call t.IsZero() to check for error.
// IMPORTANT: correct time format is
// "2014-11-19T00:00:00-08:00" for a specific time zone, or
// "2006-01-02T15:04:05Z" for UTC.
// See http://www.ietf.org/rfc/rfc3339.txt
func (xp XP) As_Time(key string) (t time.Time) {

	if val, ok := xp[key]; ok && val != nil {
		t, _ = time.Parse(time.RFC3339, fmt.Sprint(val))
	}
	return
}

// Gets parameter as time duration.
func (xp XP) As_TimeDuration(key string) (d time.Duration) {

	d = time.Duration(xp.As_Int64(key))
	return
}

// Gets parameter as time.Duration array.
func (xp XP) As_ArrayDuration(key string) (data []time.Duration) {

	if val, ok := xp[key]; ok && val != nil {
		if data, ok = val.([]time.Duration); ok {
			return
		}
	}
	return
}
