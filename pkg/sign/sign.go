package sign

import (
	"crypto/md5"
	"fmt"
	"github.com/dalebao/user_control/pkg"
)

func ValidateSign(guard, sKey string) bool {
	secret := setting.SignSecret
	key := []byte(guard + secret)
	md5String := fmt.Sprintf("%x",md5.Sum(key))
	if sKey == md5String {
		return true
	}
	return false
}
