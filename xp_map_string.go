package xparam

import "fmt"

//------------------------------------------------------------
// Xparam access for string maps
//------------------------------------------------------------

// Gets parameter as map of string to sting.
func (xp XP) As_MapString(key string) (data map[string]string) {

	if val, ok := xp[key]; ok && val != nil {
		if amap, ok := val.(map[string]interface{}); ok {
			data = map[string]string{}
			for k, v := range amap {
				if v != nil {
					data[k] = fmt.Sprint(v)
				}
			}
			return
		}
	}
	return
}
