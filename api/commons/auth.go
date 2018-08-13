/**
 * @Author codeAC
 * @Time: 2018/8/13 9:58
 * @Description
 */
package commons

import (
	"net/http"
	"project/project/videoServer/api/session"

	"project/project/videoServer/api/defs"
)

var HEADER_FILELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

//check session
func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get((HEADER_FILELD_SESSION))
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

//check user
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
