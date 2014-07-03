package xparam

import (
	"strconv"
	"strings"
	"time"
)

//------------------------------------------------------------
// Create XP from meta/data
//------------------------------------------------------------

// Creates new XP from supplied data and its metadata.
func New_FromMetaData(meta, data map[string]string) (xp XP, err error) {
	xp = XP{}

	for key, val := range data {
		if err = xp.add_MetaData(key, val, meta[key]); err != nil {
			return
		}
	}
	return
}

//------------------------------------------------------------
// Xparam access for meta data
//------------------------------------------------------------

// Adds single data entry according to its meta.
func (xp XP) add_MetaData(key, data string, meta string) (err error) {

	switch meta {

	case "[]int":
		if xp[key], err = asArray_Int(data); err != nil {
			return
		}

	case "[]duration":
		if xp[key], err = asArray_Duration(data); err != nil {
			return
		}

	default:
		xp[key] = data
	}
	return
}

//------------------------------------------------------------
// Implementation
//------------------------------------------------------------

// Parse string as array on integers.
func asArray_Int(data string) (vals []int, err error) {

	ss := strings.Split(data, ",")
	vals = make([]int, len(ss))
	var v int

	for i, s := range ss {
		s = strings.TrimSpace(s)
		if v, err = strconv.Atoi(s); err == nil {
			vals[i] = v
		} else {
			return
		}
	}
	return
}

// Parse string as array of time Durations.
func asArray_Duration(data string) (vals []time.Duration, err error) {

	ss := strings.Split(data, ",")
	vals = make([]time.Duration, len(ss))
	var v time.Duration

	for i, s := range ss {
		s = strings.TrimSpace(s)
		if v, err = StringToDuration(s); err == nil {
			vals[i] = v
		} else {
			return
		}
	}
	return
}
