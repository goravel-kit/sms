package sms

import (
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"

	"github.com/goravel-kit/sms/contracts"
)

type SMS struct {
	config map[string]string
	driver contracts.SMS
}

func NewSMS(config config.Config) *SMS {
	var info map[string]string

	driver := config.Get("sms.driver")
	switch driver {
	case "aliyun":
		info = cast.ToStringMapString(config.Get("sms.aliyun"))
	case "tencent":
		info = cast.ToStringMapString(facades.Config().Get("sms.tencent"))
	}

	return &SMS{
		config: info,
		driver: &Tencent{},
	}
}

func (r *SMS) Send(phone string, message contracts.Message) error {
	return r.driver.Send(phone, message, r.config)
}
