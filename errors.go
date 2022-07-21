package freshchat

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	// ErrNilArguments is returned when the arguments are nil.
	ErrNilArguments = errors.New("nil arguments")
)

// Error implements the error interface.
func (r *ResponseFailed) Error() string {
	return fmt.Sprintf(
		"error response from Freshchat, success:%t code:%d message:%s",
		r.Success,
		r.ErrorCode,
		r.ErrorMessage,
	)
}
