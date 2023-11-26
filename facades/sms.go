package facades

import (
	"log"

	"github.com/goravel-kit/sms"
	"github.com/goravel-kit/sms/contracts"
)

func SMS() contracts.SMS {
	instance, err := sms.App.Make(sms.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.SMS)
}
