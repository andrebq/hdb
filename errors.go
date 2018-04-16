package hdb

type (
	// Error type represents static error descriptors
	Error string
)

const (
	// NotADescriptor is returned when a given Fd isn't valid
	NotADescriptor = Error("not a valid descriptor")
)

func (e Error) Error() string {
	return string(e)
}
