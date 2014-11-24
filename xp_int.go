package xparam

import (
	"fmt"
	"strconv"
)

//------------------------------------------------------------
// Xparam access for int
//------------------------------------------------------------

// Gets parameter as int.
func (xp XP) As_Int(key string) (n int) {

	if val, ok := xp[key]; ok && val != nil {
		if n, ok = val.(int); ok {
			return
		}
		if f, ok := val.(float64); ok {
			n = int(f) // this may kill precision !
			return

		} else {
			n, _ = strconv.Atoi(fmt.Sprint(val))
		}
	}
	return
}

// Gets parameter as int64.
func (xp XP) As_Int64(key string) (n int64) {

	if val, ok := xp[key]; ok && val != nil {
		if n, ok = val.(int64); ok {
			return
		}
		if f, ok := val.(float64); ok {
			n = int64(f)
			return

		} else {
			ni, _ := strconv.Atoi(fmt.Sprint(val))
			n = int64(ni)
		}
	}
	return
}

// Gets parameter as int array.
func (xp XP) As_ArrayInt(key string) (data []int) {

	if val, ok := xp[key]; ok && val != nil {
		if data, ok = val.([]int); ok {
			return
		}
	}
	return
}
