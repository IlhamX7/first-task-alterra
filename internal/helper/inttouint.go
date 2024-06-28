package helper

import (
	"errors"
)

func IntToUint(value int) (uint, error) {
	if value < 0 {
		return 0, errors.New("cannot convert negative integer to uint")
	}
	return uint(value), nil
}
