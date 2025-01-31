package usererrors

type UseCaseError struct {
	errorDescription
}

func (e *UseCaseError) Error() string {
	return errUserMap[e.errorDescription]
}

var (
	ErrNilUserRepo = &UseCaseError{errUserNilRepo}
)