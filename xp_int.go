package xparam

//------------------------------------------------------------
// Xparam access for int
//------------------------------------------------------------

// Gets parameter as int array.
func (xp XP) As_ArrayInt(key string) (data []int) {

	if val, ok := xp[key]; ok && val != nil {
		if data, ok = val.([]int); ok {
			return
		}
	}
	return
}
