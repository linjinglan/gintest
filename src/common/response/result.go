package response

import (
	"github.com/gin-gonic/gin"
	"github.com/linjinglan/gittest/src/common/constant"
	"net/http"
)

type ResultVO struct {
	Code constant.ResponseCode `json:"code"`
	Msg constant.ResponseMsg `json:"msg"`
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}

func Success(ctx *gin.Context, code constant.ResponseCode, msg constant.ResponseMsg, data interface{})  {
	resp := &ResultVO{Code: code, Msg:msg, Success: true, Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func Failure(ctx *gin.Context, code constant.ResponseCode, msg constant.ResponseMsg, data interface{}) {
	resp := &ResultVO{Code: code, Msg: msg, Success: false, Data: data}
	ctx.JSON(http.StatusInternalServerError, resp)
}