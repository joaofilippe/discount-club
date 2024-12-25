package discounterrors

type UseCaseError struct {
	errorDescription
}

func (e *UseCaseError) Error() string {
	return errDiscountMap[e.errorDescription]
}

var (
	ErrNilDiscountRepo      = &UseCaseError{errDiscountNilRepo}
	ErrCodeProvidedOnCreate = &UseCaseError{errDiscountCodeProvidedOnCreate}
	ErrNoDiscountProvided   = &UseCaseError{errDiscountNoProvided}
)

type RepositoryError struct {
	errorDescription
}

func (e *RepositoryError) Error() string {
	return errDiscountMap[e.errorDescription]
}