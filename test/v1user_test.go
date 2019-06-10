package test

import (
	"encoding/xml"
	"fmt"
	"github.com/dalebao/user_control/logic"
	"github.com/dalebao/user_control/pkg/sign"
	"github.com/dalebao/user_control/pkg/sms"
	"testing"
)

func TestGetVerifyCodeForRAndL(t *testing.T) {
	err := logic.GetVerifyCodeForRAndL("wyche", "17681884921")
	fmt.Println(err)
}

func TestXml(t *testing.T) {
	xmlString := `<?xml version="1.0"  ="utf-8" ?><returnsms>
 <returnstatus>Success</returnstatus>
 <message>ok</message>
 <remainpoint>24214</remainpoint>
 <taskID>17685198</taskID>
 <successCounts>1</successCounts></returnsms>`
	smsResult := smsSetting.SmsResult{}
	err := xml.Unmarshal([]byte(xmlString), &smsResult)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(smsResult.ReturnStatus)
}

func TestSign(t *testing.T){
	b := sign.ValidateSign("wyche","ad69858b32e652bc8f533076349f743b")
	fmt.Println(b)
}

