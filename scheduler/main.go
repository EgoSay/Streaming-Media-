/**
 * @Author codeAC
 * @Time: 2018/8/13 20:34
 * @Description
 */

package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"project/project/videoServer/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}
func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}
