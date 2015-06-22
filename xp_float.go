package xparam

import (
	"fmt"
	"strconv"
)

//------------------------------------------------------------
// Xparam access for float
//------------------------------------------------------------

// Gets parameter as float64.
func (xp XP) As_Float64(key string) (f float64) {

	if val, ok := xp[key]; ok && val != nil {
		if f, ok = val.(float64); ok {
			return
		}
		if n, ok := val.(int); ok {
			f = float64(n)
			return

		} else {
			f, _ = strconv.ParseFloat(fmt.Sprint(val), 64)
		}
	}
	return
}

// Gets parameter as float64 array.
func (xp XP) As_ArrayFloat64(key string) (data []float64) {

	if val, ok := xp[key]; ok && val != nil {
		if data, ok = val.([]float64); ok {
			return
		}
	}
	return
}
