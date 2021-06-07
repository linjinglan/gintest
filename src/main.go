package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linjinglan/gittest/src/common/route"
)

func main() {
	r := gin.New()
	r = route.PathRoute(r)
	r.Run(":8000")
}