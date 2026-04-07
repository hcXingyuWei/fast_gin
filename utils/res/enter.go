package res

import (
	"fast_gin/utils/vaildate"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Ok(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: 0, Data: data, Msg: msg})
}

func OkWithData(data any, c *gin.Context) {
	Ok(data, "成功", c)
}

func OkWithList(list any, count int64, c *gin.Context) {
	Ok(map[string]any{
		"list":  list,
		"count": count,
	}, "成功", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Ok(gin.H{}, msg, c)
}

func Fail(code int, msg string, c *gin.Context) {
	c.JSON(200, Response{Code: code, Data: gin.H{}, Msg: msg})
}

func FailWithMig(msg string, c *gin.Context) {
	Fail(7, msg, c)
}

func FailWithError(error error, c *gin.Context) {
	msg := vaildate.ValidateError(error)
	Fail(7, msg, c)
}
