package routers

import (
	"github.com/dalebao/user_control/pkg"
	"github.com/dalebao/user_control/pkg/middleware"
	"github.com/dalebao/user_control/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(middleware.GuardMiddleware())  //验证guard
	r.Use(middleware.ValidateSignSKey()) //验证签名

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("api/v1")

	{
		//登录或注册获取验证码
		apiv1.POST("/registerOrLoginGetVerifyCode", v1.RegisterLoginGetVerifyCode)
		//使用验证码登录或注册
		apiv1.POST("/registerOrLoginWithVerifyCode", v1.RegisterLoginWithVerifyCode)
		//获取用户列表
		apiv1.GET("/users", v1.GetUsers)
		//注册
		apiv1.POST("/register", v1.Register)
		//登录
		apiv1.POST("/login", v1.Login)
		//验证用户token
		apiv1.POST("/check_u_token", v1.CheckUToken)
		//验证临时token
		apiv1.POST("/check_t_token", v1.CheckTToken)
	}
	return r
}
