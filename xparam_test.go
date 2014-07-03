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
