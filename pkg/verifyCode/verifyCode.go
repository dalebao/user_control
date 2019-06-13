package verifyCode

import (
	"errors"
	"fmt"
	"github.com/dalebao/user_control/pkg/redis"
	"math/rand"
	"time"
)

type VerifyCode struct {
	Mobile string
	Action string
	Guard  string
	Code   string
}

func (verifyCode *VerifyCode) GenerateVerifyCode() {
	verifyCode.Code = fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

func (verifyCode *VerifyCode) SaveVerifyCode() {
	key := verifyCode.Action + "_" + verifyCode.Guard + "_" + verifyCode.Mobile
	redis.Client.Set(key, verifyCode.Code, time.Minute*5)
}

func (verifyCode *VerifyCode) ValidateVerifyCode() error {
	key := verifyCode.Action + "_" + verifyCode.Guard + "_" + verifyCode.Mobile
	rKey, _ := redis.Client.Get(key).Result()
	fmt.Println(rKey)
	if rKey == "" {
		return errors.New("验证码已过期")
	}

	if rKey != verifyCode.Code {
		return errors.New("验证码输入错误")
	}
	return nil
}
