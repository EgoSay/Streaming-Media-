/**
 * @Author codeAC
 * @Time: 2018/8/13 13:16
 * @Description
 */

package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many Request")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	//流传输
	router.GET("/videos/:vid-id", streamHandler)
	//文件上传
	router.POST("/upload/:vid-id", uploadHandler)
	//测试文件上传
	router.GET("/test", testPageHandler)
	return router
}
func main() {
	router := RegisterHandlers()
	handler := NewMiddleWareHandler(router, 2)
	http.ListenAndServe(":9000", handler)
}
