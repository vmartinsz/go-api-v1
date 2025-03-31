package contracts

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type PostgresIface interface {
	Client() *gorm.DB
	Close() error
}

type Logger interface {
	Infof(format string, args ...zap.Field)
	Errorf(format string, args ...zap.Field)
}
