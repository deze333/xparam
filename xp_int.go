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
		} else {
			str := fmt.Sprint(val)
			n, _ = strconv.Atoi(str)
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
