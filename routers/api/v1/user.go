package v1

import (
	"errors"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/dalebao/user_control/logic"
	"github.com/dalebao/user_control/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	maps := make(map[string]interface{})

	name := c.Query("name")

	pageNum := 1
	if pageNumO := c.Query("page"); pageNumO != "" {
		pageNum = com.StrTo(pageNumO).MustInt()
	}
	pageSize := 10
	if pageSizeO := c.Query("page_size"); pageSizeO != "" {
		pageSize = com.StrTo(pageSizeO).MustInt()
	}

	if name != "" {
		maps["name"] = name
	}
	data := make(map[string]interface{})

	data["total"], data["user_list"] = models.GetUsers(pageNum, pageSize, maps)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func Register(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")
	passwordConfirm := c.PostForm("password_confirm")
	mobile := c.PostForm("mobile")
	guard := c.PostForm("guard")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("账号不能为空")
	valid.Required(mobile, "mobile").Message("手机号不能为空")
	valid.Required(password, "password").Message("密码不能为空")
	valid.Required(passwordConfirm, "password_confirm").Message("确认密码不能为空")

	var err error
	res := make(map[string]interface{})

	if !valid.HasErrors() {
		if password == passwordConfirm {
			res, err = logic.CreateUser(name, password, mobile, guard)
		} else {
			valid.SetError("password", "两次输入的密码不同")
		}
	}

	if valid.HasErrors() {
		err = errors.New(valid.Errors[0].Message)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  err.Error(),
			"data": res,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  err,
		"data": res,
	})

}

func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	guard := c.PostForm("guard")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("账号不能为空")
	valid.Required(password, "password").Message("密码不能为空")

	var err error
	res := make(map[string]interface{})

	if valid.HasErrors() {
		err = errors.New(valid.Errors[0].Message)
	}

	res, err = logic.Login(name, password, guard)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  err.Error(),
			"data": res,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  err,
		"data": res,
	})
}

func CheckUToken(c *gin.Context) {
	uToken := c.PostForm("uToken")

	valid := validation.Validation{}
	valid.Required(uToken, "uToken").Message("用户令牌必填")

	var err error
	res := make(map[string]interface{})

	if valid.HasErrors() {
		err = errors.New(valid.Errors[0].Message)
	}

	res, err = logic.CheckUToken(uToken)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  err.Error(),
			"data": res,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  err,
		"data": res,
	})
}

func CheckTToken(c *gin.Context) {
	tToken := c.PostForm("tToken")

	valid := validation.Validation{}
	valid.Required(tToken, "tToken").Message("用户临时令牌必填")

	var err error
	res := make(map[string]interface{})

	if valid.HasErrors() {
		err = errors.New(valid.Errors[0].Message)
	}

	res, err = logic.CheckTToken(tToken)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  err.Error(),
			"data": res,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  err,
		"data": res,
	})
}

/**
登录/注册 获取验证码
 */
func RegisterLoginGetVerifyCode(c *gin.Context) {
	mobile := c.PostForm("mobile")
	guard := c.PostForm("guard")


	valid := validation.Validation{}
	valid.Required(mobile, "mobile").Message("手机号必填")
	valid.Required(guard, "guard").Message("门卫必填")


	err := logic.GetVerifyCodeForRAndL(guard, mobile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  err,
		"data": "",
	})

}
