package smsSetting

import (
	"encoding/xml"
	"github.com/dalebao/user_control/pkg"
	"github.com/go-ini/ini"
	"log"
)

type SmsConfig struct {
	UserId   string
	Account  string
	Password string
	RContent string
	LContent string
}

func (smsConfig *SmsConfig) LoadConfig(section string) {
	Cfg, err := ini.Load(setting.Dir + "/conf/app.ini")

	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	sec, err := Cfg.GetSection(section)
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	smsConfig.UserId = sec.Key("USERID").String()
	smsConfig.Account = sec.Key("ACCOUNT").String()
	smsConfig.Password = sec.Key("PASSWORD").String()
	smsConfig.RContent = sec.Key("RCONTENT").String()
	smsConfig.LContent = sec.Key("LCONTENT").String()
}

type SmsResult struct {
	ReturnSms xml.Name `xml:"returnsms"`
	ReturnStatus string `xml:"returnstatus"`
	Message string `xml:"message"`
	RemainPoint string `xml:"remainpoint"`
	TaskID string `xml:"taskID"`
	SuccessCounts string `xml:"successCounts"`
}
