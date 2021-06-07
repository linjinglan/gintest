package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/linjinglan/gittest/src/common/constant"
	"github.com/linjinglan/gittest/src/common/response"
)

type UserController struct {

}

func UserRegister(userGrp *gin.RouterGroup) {
	userController := &UserController{}
	userGrp.Use().POST("/findUser", userController.findUser)
	userGrp.Use().POST("/get/user", userController.getUser)
}

func (c UserController) findUser(ctx *gin.Context)  {
	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, gin.H{"userName": "月影"})
}

func (c UserController) getUser(ctx *gin.Context) {
	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, gin.H{"getuser": "hello"})
}