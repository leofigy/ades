package contracts

import "gorm.io/gorm"

type Provider interface {
	GetDB() (*gorm.DB, error)
}
