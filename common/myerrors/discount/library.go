package discounterrors


type errorDescription int

const (
	errDiscountNilRepo = iota
	errDiscountCodeProvidedOnCreate
	errDiscountNoProvided
)

var errDiscountMap = map[errorDescription]string{
	errDiscountNilRepo:              "discount repository is nil",
	errDiscountCodeProvidedOnCreate: "code should not be provided on create",
	errDiscountNoProvided:           "no discount provided",
}
