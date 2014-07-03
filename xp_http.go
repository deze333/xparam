package xparam

import "net/http"

//------------------------------------------------------------
// Xparam HTTP request related values
//------------------------------------------------------------

// Gets http.Request.
func (xp XP) Get_HttpRequest() (req *http.Request) {

	if val, ok := xp["_httpReq"]; ok && val != nil {
		if req, ok = val.(*http.Request); ok {
			return
		}
	}
	return
}

// Gets http.ResponseWriter.
func (xp XP) Get_HttpResponseWriter() (rw http.ResponseWriter) {

	if val, ok := xp["_httpResWriter"]; ok && val != nil {
		if rw, ok = val.(http.ResponseWriter); ok {
			return
		}
	}
	return
}

// Gets session values.
func (xp XP) Get_SessionValues() (vals map[string]string) {

	if val, ok := xp["_session"]; ok && val != nil {
		if vals, ok = val.(map[string]string); ok {
			return
		}
	}
	return
}
