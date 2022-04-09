package brainfuck

import "errors"

var (
	OperationExistsError = errors.New("operation already exists")
)

type Operation func(b *Brainfuck) error

type CustomOperation struct {
	Character rune
	Operation Operation
}

type Options func(b *Brainfuck) error
