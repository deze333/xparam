package xparam

import (
	"fmt"
	"testing"
)

//------------------------------------------------------------
// Test
//------------------------------------------------------------

// Parse meta/data.
func TestParse(t *testing.T) {

	var err error
	var xp XP

	var meta map[string]string
	var data map[string]string

	// Test data
	meta = map[string]string{
		"ids": "[]int",
		"dt":  "[]duration",
	}

	data = map[string]string{
		"name": "parabits",
		"ids":  "1, 17,  100,123",
		"dt":   "25h,  24h,7h,5m",
	}

	if xp, err = New_FromMetaData(meta, data); err == nil {
		fmt.Println(xp)
	} else {
		t.Error(err)
	}
}

// Parse array of string.
func TestArrayString(t *testing.T) {

	data := map[string]interface{}{
		"txt": []interface{}{"one", "at", "a", "time"},
	}

	xp := New(data)

	arr := xp.As_ArrayString("txt")
	fmt.Println(arr)

	if len(arr) == 0 {
		t.Error("Failed to convert to array of strings")
		return
	}
}

// Parse time duration.
func TestTimeDuration(t *testing.T) {

	data := map[string]interface{}{
		"durA": "3600000000000",
		"durB": 3600000000000,
		"durC": float64(3600000000000),
	}

	xp := New(data)

	// From string
	dur := xp.As_TimeDuration("durA")
	fmt.Println(dur)

	if dur != 3600000000000 {
		t.Error("Failed to convert time duration from string")
		return
	}

	// From int
	dur = xp.As_TimeDuration("durB")
	fmt.Println(dur)

	if dur != 3600000000000 {
		t.Error("Failed to convert time duration from int")
		return
	}

	// From float
	dur = xp.As_TimeDuration("durC")
	fmt.Println(dur)

	if dur != 3600000000000 {
		t.Error("Failed to convert time duration from float")
		return
	}
}
