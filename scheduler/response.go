/**
 * @Author codeAC
 * @Time: 2018/8/13 20:34
 * @Description
 */
package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, sc int, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
