/**
 * Created by angelina on 2017/7/29.
 */

package yeeSubmail

import (
	"encoding/json"
	"errors"
)

func submailSendSmsCode(phoneNum, code string) error {
	config := Config{
		AppId:    "submail.AppId",
		AppKey:   "submail.AppKey",
		SignType: "md5",
	}
	mXSend := CreateMessageXSend(phoneNum, "x25na")
	mXSend.AddVar("code", code)
	result := mXSend.Run(config)
	res := &SubmailResponse{}
	if err := json.Unmarshal([]byte(result), res); err != nil {
		return err
	}
	if res.Status != "success" {
		return errors.New("error")
	}
	return nil
}
