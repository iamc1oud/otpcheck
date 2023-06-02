package store

import (
	"errors"

	"github.com/iamc1oud/otpcheck/pkg/models"
)

var ErrNotExist = errors.New("the OTP does not exist")

type Store interface {
	Set(namespace, id string, otp models.OTP) (models.OTP, error)

	// SetAddress sets (updates) the address on an existing OTP.
	SetAddress(namespace, id, address string) error

	// Check checks the attempt count and TTL duration against an ID.
	// Passing counter=true increments the attempt counter.
	Check(namespace, id string, counter bool) (models.OTP, error)

	// Close closes an OTP and marks it as done (verified).
	// After this, the OTP has to expire after a TTL or be deleted.
	Close(namespace, id string) error

	// Delete deletes the OTP saved against a given ID.
	Delete(namespace, id string) error

	// Ping checks if store is reachable
	Ping() error
}
