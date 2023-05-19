package DBReader

import "io"

type DBReader interface {
	Parse(io.Reader) error
	ConvertPP() ([]byte, error)
	WriteToAnotherFormat([]byte) error
}
