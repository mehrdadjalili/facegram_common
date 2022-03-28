package derrors

import (
	"github.com/mehrdadjalili/facegram_common/pkg/translator/messages"
)

func InternalError() error {
	return New(StatusInternalServerError, messages.InternalServerError)
}

func BadRequest() error {
	return New(StatusBadRequest, messages.BadRequest)
}

func Unauthorized() error {
	return New(StatusUnauthorized, messages.Unauthorized)
}

func PaymentRequired() error {
	return New(StatusPaymentRequired, messages.PaymentRequired)
}

func Forbidden() error {
	return New(StatusForbidden, messages.Forbidden)
}

func NotAcceptable() error {
	return New(StatusNotAcceptable, messages.NotAcceptable)
}

func UnsupportedMediaType() error {
	return New(StatusUnsupportedMediaType, messages.UnsupportedMediaType)
}

func Locked() error {
	return New(StatusLocked, messages.Locked)
}

func UpgradeRequired() error {
	return New(StatusUpgradeRequired, messages.UpgradeRequired)
}

func TooManyRequests() error {
	return New(StatusLocked, messages.Locked)
}
