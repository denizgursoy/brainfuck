package brainfuck

import "errors"

var (
	OperationExistsError   = errors.New("operation already exists")
	OperationNilError      = errors.New("operation can not be nil")
	CommandReaderNilError  = errors.New("command reader can not be empty")
	InputReaderNilError    = errors.New("input reader can not be empty")
	OutputWriterNilError   = errors.New("output writer can not be empty")
	LoopEndInvalidError    = errors.New("invalid loop end because it did not start")
	LoopEndIsNotFoundError = errors.New("end of the loop is not found while executing one command at a time ")
)

// Operation is a function which is related to a Command
// and executed when the related Command is pointed
type Operation func(b *Brainfuck) error

// Option is supplied by user so that it can customize a brainfuck
type Option func(b *Brainfuck) error
