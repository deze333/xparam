package xparam

//------------------------------------------------------------
// Xparam
//------------------------------------------------------------

func New(vals map[string]interface{}) XP {
	return vals
}

//------------------------------------------------------------
// Xparam access for inner xparams
//------------------------------------------------------------

// Gets parameter as array of xparams.
func (xp XP) As_XP(key string) (data XP) {

	if val, ok := xp[key]; ok && val != nil {
		if axp, ok := val.(map[string]interface{}); ok {
			data = axp
		}
	}
	return
}

// Gets parameter as array of xparams.
func (xp XP) As_ArrayXP(key string) (data []XP) {

	if val, ok := xp[key]; ok && val != nil {
		if arr, ok := val.([]interface{}); ok {
			data = []XP{}
			for _, obj := range arr {
				if axp, ok := obj.(map[string]interface{}); ok {
					data = append(data, axp)
				}
			}
		}
	}
	return
}
