package discounterrors

type errorDescription int

const (
	errDiscountNilRepo = iota
	errDiscountCodeProvidedOnCreate
	errDiscountNoProvided
	errNoDiscountIDProvided
	errDiscountNotFound
	errInvalidDiscountID
)

var errDiscountMap = map[errorDescription]string{
	errDiscountNilRepo:              "discount repository is nil",
	errDiscountCodeProvidedOnCreate: "code should not be provided on create",
	errDiscountNoProvided:           "no discount provided",
	errNoDiscountIDProvided:         "no discount ID provided",
	errDiscountNotFound:             "discount not found",
	errInvalidDiscountID:            "invalid discount ID",
}
