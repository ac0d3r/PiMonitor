package handler

import (
	"pimonitor/database"
	"pimonitor/handler/base"
	"pimonitor/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserLogin 用户登陆
func UserLogin(c *gin.Context) {
	var param = &struct {
		Username string `json:"username" form:"username" validate:"required"`
		Password string `json:"password" form:"password" validate:"required"`
	}{}
	if err := c.ShouldBind(param); err != nil {
		base.ReturnError(c, 1001, err)
		return
	}

	u, err := database.FindUserByUserName(param.Username)
	if err != nil || u.ID == 0 {
		base.ReturnError(c, 2002, err)
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(param.Password)); err != nil {
		base.ReturnError(c, 2002, err)
		return
	}

	// 生成token
	token := util.RandomString(32)
	// 更新token
	if err := database.UpdateUserToken(u.ID, token); err != nil {
		base.ReturnError(c, 1002, err)
		return
	}

	base.ReturnJSON(c, map[string]interface{}{
		"id":       u.ID,
		"username": u.Username,
		"token":    token,
	})
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	u, err := database.FindUserByID(base.GetUserID(c))
	if err != nil || u.ID == 0 {
		base.ReturnError(c, 1002, err)
		return
	}
	// remove token
	if err := database.UpdateUserToken(u.ID, ""); err != nil {
		base.ReturnError(c, 1002, err)
		return
	}

	base.ReturnJSON(c, nil)
}

// UserUpdateInfomation 用户信息更新
func UserUpdateInfomation(c *gin.Context) {
	var param = &struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}{}
	if err := c.ShouldBind(param); err != nil {
		base.ReturnError(c, 1001, err)
		return
	}

	u, err := database.FindUserByID(base.GetUserID(c))
	if err != nil || u.ID == 0 {
		base.ReturnError(c, 1002, err)
		return
	}

	if param.Username != "" {
		u.Username = param.Username
	}
	if param.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
		if err != nil {
			base.ReturnError(c, 1001, err)
			return
		}
		u.Password = string(hashed)
	}

	if err := database.UpdateUser(u); err != nil {
		base.ReturnError(c, 1002, err)
		return
	}

	base.ReturnJSON(c, nil)
}
