package utils

import (
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/logger"
	"go.uber.org/zap"
)

// Log err
func OnError(err error) {
	logger.LoadLogger()
	zap.L().Error(err.Error())
}
