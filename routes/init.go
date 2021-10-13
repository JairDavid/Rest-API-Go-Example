package routes

import "github.com/gin-gonic/gin"

func StartRouting(server *gin.Engine) {

	api := *server.Group("/api/v1")

	CategoryRouter(&api)
	CustomerRouter(&api)
	ProductRouter(&api)
	CustomerProductRouter(&api)

}
