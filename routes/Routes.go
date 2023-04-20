package routes

import (
	Nasabah "go_rest/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	nApi := r.Group("/api")
	{
		nApi.GET("nasabah", Nasabah.GetAll)
		nApi.POST("nasabah", Nasabah.Create)
		nApi.GET("nasabah/:id", Nasabah.Detail)
		nApi.PUT("nasabah/:id", Nasabah.Update)
		nApi.DELETE("nasabah/:id", Nasabah.Delete)
	}
	return r
}
