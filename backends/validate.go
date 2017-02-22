package backends

import (
	"errors"
)

type RcptError error

var (
	NoSuchUser          = RcptError(errors.New("no such iser"))
	StorageNotAvailable = RcptError(errors.New("storage not available"))
	StorageTooBusy      = RcptError(errors.New("stoarge too busy"))
	StorageTimeout      = RcptError(errors.New("stoarge too busy"))
	QuotaExceeded       = RcptError(errors.New("quota exceeded"))
	UserSuspended       = RcptError(errors.New("user suspended"))
)
