package usererrors

type errorDescription int

const (
	errUserNilRepo = iota
)

var errUserMap = map[errorDescription]string{
	errUserNilRepo: "user repository is nil",
}