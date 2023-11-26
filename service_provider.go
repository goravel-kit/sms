package sms

import (
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "sms"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewSMS(app.MakeConfig()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/goravel-kit/sms", map[string]string{
		"config/sms.go": app.ConfigPath("sms.go"),
	})
}
