package xparam

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

//------------------------------------------------------------
// Time Duration conversion
//------------------------------------------------------------

// Converts string into duration, ie:
// 1d, 1h, 1m
// 30d, 77h, 99m
func StringToDuration(val string) (dt time.Duration, err error) {

	num := 0
	l := len(val)
	if l == 0 {
		err = errors.New("Cannot convert empty string to duration")
		return
	}
	num, err = strconv.Atoi(val[0 : l-1])
	if err != nil {
		return
	}

	switch val[l-1 : l] {
	case "d":
		dt = time.Duration(num) * 24 * time.Hour
	case "h":
		dt = time.Duration(num) * time.Hour
	case "m":
		dt = time.Duration(num) * time.Minute
	case "s":
		dt = time.Duration(num) * time.Second
	default:
		err = fmt.Errorf("Time duration must end with d/h/m/s: %v", val)
	}
	return
}

// Converts duration into string notation, ie:
// 1d, 1h, 1m
// The smallest interval is 1m. Smaller values replaced with 1m.
func DurationToString(dt time.Duration) string {
	h := dt / time.Hour
	hm := dt % time.Hour
	d := h / 24
	dh := dt % (24 * time.Hour)

	// If matches whole day
	if d > 0 && dh == 0 {
		return fmt.Sprintf("%dd", d)
	}

	// If matches whole hour
	if h > 0 && hm == 0 {
		return fmt.Sprintf("%dh", h)
	}

	// Otherwise make into minutes
	m := dt / time.Minute
	if m > 0 {
		return fmt.Sprintf("%dm", m)
	}

	return "1m"
}
