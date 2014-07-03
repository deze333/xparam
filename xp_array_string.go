package xparam

//------------------------------------------------------------
// Xparam access for string arrays
//------------------------------------------------------------

// Gets parameter as string array.
func (xp XP) As_ArrayString(key string) (data []string) {

	if val, ok := xp[key]; ok && val != nil {
		if data, ok = val.([]string); ok {
			return
		}
	}
	return
}
