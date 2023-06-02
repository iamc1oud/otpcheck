package store

import "errors"

var ErrNotExist = errors.New("the OTP does not exist")

type Store interface {
	Set(namespace, id string, otp models.OTP) (models.OTP, error)
}
