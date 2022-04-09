package brainfuck

import "errors"

var (
	OperationExistsError  = errors.New("operation already exists")
	OperationNilError     = errors.New("operation can not be nil")
	CommandReaderNilError = errors.New("command reader can not be empty")
	InputReaderNilError   = errors.New("input reader can not be empty")
	OutputWriterNilError  = errors.New("output writer can not be empty")
)

type Operation func(b *Brainfuck) error

type Option func(b *Brainfuck) error
