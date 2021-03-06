package xparam

import "fmt"

//------------------------------------------------------------
// Xparam access for bool
//------------------------------------------------------------

// Gets parameter as bool.
// Can be either pure JSON bool or "true" string.
// All other stings including empty will be treated as false.
func (xp XP) As_Bool(key string) (b bool) {

	if val, ok := xp[key]; ok && val != nil {
		if b, ok = val.(bool); ok {
			return
		} else {
			str := fmt.Sprint(val)
			if str == "true" {
				return true
			} else if str == "false" {
				return false
			}
			return false
		}
	}
	return
}
