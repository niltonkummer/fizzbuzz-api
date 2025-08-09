package model

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}

var (
	ErrNoRequestsFound = &Error{
		Code:    "no_requests_found",
		Message: "No requests found in the statistics",
	}
)
