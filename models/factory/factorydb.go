package factory

import "gorm.io/gorm"

type FactoryDB interface {
	GetDB(usuario int) *gorm.DB
}
