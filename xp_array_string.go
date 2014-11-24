package xparam

import "fmt"

//------------------------------------------------------------
// Xparam access for string arrays
//------------------------------------------------------------

// Gets parameter as string array.
func (xp XP) As_ArrayString(key string) (data []string) {

	if val, ok := xp[key]; ok && val != nil {
		if raw, ok := val.([]interface{}); ok {
			data = make([]string, len(raw))
			for i, o := range raw {
				data[i] = fmt.Sprint(o)
			}
			return
		}
	}
	return
}
