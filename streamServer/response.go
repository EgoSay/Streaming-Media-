/**
 * @Author codeAC
 * @Time: 2018/8/13 13:42
 * @Description
 */
package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
