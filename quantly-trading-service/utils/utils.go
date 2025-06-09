package utils

import (
	"fmt"
)

func OnError(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("%v %w", msg, err)
	}
	return nil
}
