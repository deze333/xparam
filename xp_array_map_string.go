package xparam

import "fmt"

//------------------------------------------------------------
// Xparam access for array of string maps
//------------------------------------------------------------

// Gets parameter as array of string maps.
func (xp XP) As_ArrayMapString(key string, fields ...string) (data []map[string]string) {

	if val, ok := xp[key]; ok && val != nil {
		if arr, ok := val.([]interface{}); ok {
			data = []map[string]string{}

			for _, obj := range arr {
				if inmap, ok := obj.(map[string]interface{}); ok {
					outmap := map[string]string{}
					for k, v := range inmap {
						isPresent := false
						for _, field := range fields {
							if k == field {
								isPresent = true
								break
							}
						}
						if isPresent && v != nil {
							outmap[k] = fmt.Sprint(v)
						}
					}
					if len(outmap) != 0 {
						data = append(data, outmap)
					}
				}
			}
			return
		}
	}
	return
}
