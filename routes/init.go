package routes

import "github.com/gin-gonic/gin"

func StartRouting(server *gin.Engine) {

	//Global group
	api := *server.Group("/api/v1")

	//Subgroups
	CategoryRouter(&api)
	CustomerRouter(&api)
	ProductRouter(&api)
}
