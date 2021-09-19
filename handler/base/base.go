package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func ReturnJSON(c *gin.Context, data interface{}) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1000, "msg": "系统内部错误"})
		}
	}()
	c.JSON(http.StatusOK, Result{
		Data: data,
	})
}

func ReturnError(c *gin.Context, code int, err ...error) {
	msg, has := errorMap[code]
	if !has {
		msg = "unknown error code"
	}

	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
	})
}

func ReturnUnauthorized(c *gin.Context, code int, err ...error) {
	msg, has := errorMap[code]
	if !has {
		msg = "unknown error code"
	}

	c.JSON(http.StatusUnauthorized, Result{
		Code: code,
		Msg:  msg,
	})
}

var errorMap = map[int]string{
	1000: "服务器内部错误",
	1001: "请求参数错误",
	1002: "数据库错误",
	2000: "未授权访问",
	2001: "登陆状态已过期",
	2002: "用户名或密码错误",
}

func GetUserID(c *gin.Context) uint {
	itr, has := c.Get("auth.user_id")
	if !has {
		ReturnUnauthorized(c, 2000)
		c.Abort()
	}
	return itr.(uint)
}
