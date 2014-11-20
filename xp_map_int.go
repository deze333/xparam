package xparam

import (
	"fmt"
	"strconv"
)

//------------------------------------------------------------
// Xparam access for integer maps
//------------------------------------------------------------

// Gets parameter as map of string to integers.
func (xp XP) As_MapInt(key string) (data map[string]int64) {

	if val, ok := xp[key]; ok && val != nil {
		if amap, ok := val.(map[string]interface{}); ok {
			data = map[string]int64{}
			for k, v := range amap {
				if v != nil {
					data[k], _ = strconv.ParseInt(fmt.Sprint(v), 10, 16)
				}
			}
			return
		}
	}
	return
}
