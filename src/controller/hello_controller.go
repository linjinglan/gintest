package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/linjinglan/gittest/src/common/constant"
	"github.com/linjinglan/gittest/src/common/response"
)

type HelloController struct {

}

func HelloRegister(userGrp *gin.RouterGroup) {
	helloController := &HelloController{}
	userGrp.Use().GET("/hello/get", helloController.get)
}

func (c HelloController) get(ctx *gin.Context)  {
	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, gin.H{"userName": "hello get"})
}