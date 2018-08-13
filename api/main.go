package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"project/project/videoServer/api/commons"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	commons.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

//用户注册
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	//创建用户
	router.POST("/user", commons.CreateUser)
	//用户登录
	router.POST("/user/:user_name", commons.Login)
	return router
}

func main() {
	r := RegisterHandlers()
	handler := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", handler)
}
