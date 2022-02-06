package repo

import (
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Init(db *gorm.DB) error {
	dbInstance = db

	// Register model types
	if err := dbInstance.AutoMigrate(&model.AuthenticatorModel{}); err != nil {
		return errs.Handle(nil, err)
	}
	if err := dbInstance.AutoMigrate(&model.Authenticator{}); err != nil {
		return errs.Handle(nil, err)
	}
	return nil
}
