package route

import (
	"github.com/gin-gonic/gin"
	"github.com/linjinglan/gittest/src/controller"
)

func PathRoute(r *gin.Engine) *gin.Engine {
	rootPath := r.Group("/api")
	{
		controller.UserRegister(rootPath)
		controller.HelloRegister(rootPath)
	}

	return r
}