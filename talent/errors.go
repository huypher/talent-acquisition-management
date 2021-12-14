package talent

func NewErrTalentNotFound(msg string) *ErrTalentNotFound {
	return &ErrTalentNotFound{message: msg}
}

type ErrTalentNotFound struct {
	message string
}

func (e *ErrTalentNotFound) Error() string {
	return e.message
}
