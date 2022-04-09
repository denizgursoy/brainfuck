package brainfuck

import "errors"

var (
	OperationExistsError = errors.New("operation already exists")
	OperationNilError    = errors.New("operation can not be nil")
)

type Operation func(b *Brainfuck) error

type CustomOperation struct {
	Character rune
	Operation Operation
}

type Options func(b *Brainfuck) error
