package health

// @Enum
type Status uint8

const (
	StatusUnknown Status = iota
	StatusServing
	StatusNotServing
)
