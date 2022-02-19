package derrors

import (
	"github.com/mehrdadjalili/facegram_common/pkg/logger"
	"github.com/mehrdadjalili/facegram_common/pkg/translator/messages"
)

func InternalError() error {
	return New(KindUnexpected, messages.GeneralError)
}

func InternalErrorWithLogger(function string, err error, logger logger.Logger) error {
	return NewWithLogger(
		KindUnexpected,
		messages.GeneralError,
		logger,
		function,
		err,
	)
}
