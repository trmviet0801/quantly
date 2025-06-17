package utils

import (
	"fmt"
	"regexp"
	"strings"

	"go.uber.org/zap"
)

func OnError(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("%v %w", msg, err)
	}
	return nil
}

func OnLogError(err error, msg string) {
	if err != nil {
		zap.L().Error(msg)
	}
}

func RemoveSpecialSymbol(value string) string {
	reg := regexp.MustCompile(`[()%,]`)
	return reg.ReplaceAllString(strings.TrimSpace(value), "")
}
