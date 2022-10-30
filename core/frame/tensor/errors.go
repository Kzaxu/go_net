package tensor

import "fmt"

type SizeError struct {
	ESize []uint
	ASize []uint
}

func (e *SizeError) Error() string {
	msg := fmt.Sprintf("expect size %v, actually %v",
		e.ESize, e.ASize)
	return msg
}
