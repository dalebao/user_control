package logic

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/dalebao/user_control/models"
	"github.com/dalebao/user_control/pkg"
	"github.com/dalebao/user_control/pkg/request"
	"github.com/dalebao/user_control/pkg/sms"
	"github.com/dalebao/user_control/pkg/verifyCode"
	"strings"
)

/**
注册用户
 */
func CreateUser(name, password, mobile, guard string) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	user, err := models.CreateUser(name, password, mobile)
	if err != nil {
		return res, err
	}

	res["user"] = user
	res["u_token"], err = GenerateUToken(user.Name, user.Uuid, guard)
	return res, err
}

/**
登录注册获取验证码
 */

func GetVerifyCodeForRAndL(guard, mobile string) error {
	smsConfig := &smsSetting.SmsConfig{}
	smsConfig.LoadConfig(guard)

	action := "register"
	content := smsConfig.RContent

	user := models.FindUserByMobile(mobile)
	if user.ID > 0 {
		action = "login"
		content = smsConfig.LContent
	}

	v := &verifyCode.VerifyCode{Mobile: mobile, Guard: guard, Action: action}
	v.GenerateVerifyCode()

	content = strings.Replace(content, "verifyCode", v.Code, 1)

	params := make(map[string]string)

	params["userid"] = smsConfig.UserId
	params["account"] = smsConfig.Account
	params["password"] = smsConfig.Password
	params["content"] = content
	params["mobile"] = mobile
	params["sendtime"] = ""

	smsResultXml := request.HttpPostForm(setting.SmsUrl, params)

	smsResult := smsSetting.SmsResult{}
	err := xml.Unmarshal([]byte(smsResultXml), &smsResult)
	if err != nil {
		fmt.Printf("error: %v", err)
		return errors.New("xml 解析出错")
	}

	if (smsResult.ReturnStatus != "Success") {
		return errors.New("短信发送失败")
	}
	//保存验证码
	v.SaveVerifyCode()

	return nil
}

/**
账号密码登录
 */
func Login(name, password, guard string) (map[string]interface{}, error) {
	res := make(map[string]interface{})

	user, err := models.Login(name, password)

	if err != nil {
		return res, err
	}

	res["user"] = user
	res["token"], err = GenerateUToken(user.Name, user.Uuid, guard)
	return res, err
}

/**
验证utoken 获取tmptoken
 */
func CheckUToken(uToken string) (map[string]interface{}, error) {
	var err error
	res := make(map[string]interface{})
	claims := &Claims{}
	claims, err = ParseUToken(uToken)

	if err != nil {
		return res, err
	}

	uuid := claims.Uuid
	guard := claims.Guard

	res["user"], err = models.FindUserByUuid(uuid)
	res["guard"] = guard
	res["t_token"], err = GenerateTToken(uToken, uuid, guard)

	return res, err
}

/**
验证tmptoken 获取utoken
 */
func CheckTToken(tToken string) (map[string]interface{}, error) {
	var err error
	res := make(map[string]interface{})
	claims := &TClaims{}
	claims, err = ParseTToken(tToken)

	if err != nil {
		return res, err
	}

	uuid := claims.Uuid
	guard := claims.Guard

	user := models.User{}
	user, err = models.FindUserByUuid(uuid)
	if err != nil {
		return res, err
	}
	name := user.Name
	res["user"] = user
	res["guard"] = guard
	res["u_token"], err = GenerateUToken(name, uuid, guard)

	return res, err
}
