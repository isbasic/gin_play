package routes

import (
	"html/template"

	"github.com/isbasic/gin_play/common"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Run will start the server
// func Run() {
// 	getRoutes()
// 	router.Run(":5000")
// }

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func GetRoutes() *gin.Engine {
	router.SetFuncMap(template.FuncMap{
		"B64": common.B64Encode,
	})

	router.LoadHTMLGlob("templates/*")

	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)
	addPicRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)

	return router
}
