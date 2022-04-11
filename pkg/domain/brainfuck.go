package brainfuck

import (
	"io"
)

const (
	InitialCapacity = 1000
)

// NewBrainFuck creates a new brainfuck. Requires IoOptions pointers which has no nil
// reader or writers. A brainfuck might have custom defined operation, so it also receives
// options as much as user pleases. IoOptions fields are checked if they exist, or it returns
// CommandReaderNilError, InputReaderNilError,OutputWriterNilError errors.
func NewBrainFuck(io *IoOptions, options ...Option) (*Brainfuck, error) {
	brainfuck, err := createNewBrainFuck(io)
	if err != nil {
		return nil, err
	}
	brainfuck.addDefaultOperations()
	if err := brainfuck.registerOptions(options); err != nil {
		return nil, err
	}
	return brainfuck, nil
}

// createNewBrainFuck checks if all io options are provided and initializes slices and maps
// on brainfuck struct.
func createNewBrainFuck(io *IoOptions) (*Brainfuck, error) {

	if io.CommandReader == nil {
		return nil, CommandReaderNilError
	}

	if io.InputReader == nil {
		return nil, InputReaderNilError
	}

	if io.OutputWriter == nil {
		return nil, OutputWriterNilError
	}

	brainfuck := Brainfuck{
		operations: make(map[rune]Operation, 8),
		Data:       make([]byte, InitialCapacity),
		Commands:   make([]rune, 0),
		IoOptions:  io,
		loopStack:  make([]*Loop, 0),
	}
	return &brainfuck, nil
}

// addDefaultOperations adds 8 mandatory operations to the brainfuck
func (b *Brainfuck) addDefaultOperations() {
	_ = b.ExtendWith(CustomOperation{'+', incrementOperation})
	_ = b.ExtendWith(CustomOperation{'-', decrementOperation})
	_ = b.ExtendWith(CustomOperation{'>', shiftRightOperation})
	_ = b.ExtendWith(CustomOperation{'<', shiftLeftOperation})
	_ = b.ExtendWith(CustomOperation{'.', printOperation})
	_ = b.ExtendWith(CustomOperation{',', setFromUserInputOperation})
	_ = b.ExtendWith(CustomOperation{'[', startLoopOperation})
	_ = b.ExtendWith(CustomOperation{']', endLoopOperation})
}

// registerOptions calls options provided by the caller on NewBrainFuck.
// It allows user to customize brainfuck. In this case, it allows user to add new operations
func (b *Brainfuck) registerOptions(options []Option) error {
	for _, customOperation := range options {
		if err := customOperation(b); err != nil {
			return err
		}
	}
	return nil
}

// ExtendWith adds new operations to brainfuck.
// CustomOperation can not have characters which is already defined.
// OperationNilError is returned when operation is nil
// OperationExistsError is returned when CustomOperation's character is already defined
func (b *Brainfuck) ExtendWith(operation CustomOperation) error {

	if operation.Operation == nil {
		return OperationNilError
	}

	character := operation.Character
	if b.isCommandDefined(&character) {
		return OperationExistsError
	}

	b.operations[character] = operation.Operation
	return nil
}

// getCurrentCellValue return value of cell which data pointer is showing
func (b *Brainfuck) getCurrentCellValue() byte {
	return b.Data[b.DataPointer]
}

// Start triggers reading from CommandReader and executes the operation relating to command
// return error if executed operation returns errors.
// it executes until there is no more input to read in the happy case
func (b *Brainfuck) Start() error {

	for {
		command, inputExists := b.getCommandToExecute()
		if !inputExists {
			break
		}

		if err := b.performOperation(command); err != nil {
			return err
		}
		b.CommandPointer++
	}
	return nil
}

// getCommandToExecute calculates which command will be executed next.
// if the CommandPointer is at last item, it reads from CommandReader and executes newly read command
//f the CommandPointer is not at last item, it processes the next item.
func (b *Brainfuck) getCommandToExecute() (*rune, bool) {
	var command *rune
	if b.isCommandPointerAtLast() {
		definedCommand, inputExists := b.readCommand()
		if !inputExists {
			return nil, false
		}
		b.addNewCommand(definedCommand)
		command = definedCommand
	} else {
		command = &b.Commands[b.CommandPointer]
	}
	return command, true
}

// performOperation executes the Operation related to command
func (b *Brainfuck) performOperation(command *rune) error {
	operation := b.operations[*command]
	return operation(b)
}

// isCommandDefined checks is a command is defined before
func (b *Brainfuck) isCommandDefined(c *rune) bool {
	return b.operations[*c] != nil
}

// addNewCommand appends new command to Commands
func (b *Brainfuck) addNewCommand(command *rune) {
	b.Commands = append(b.Commands, *command)
}

// readCommand reads a rune from user command reader
// if command read from reader is not defined,
// it skips the command and reads new file
// return rune read or bool showing if there is something to read
func (b *Brainfuck) readCommand() (*rune, bool) {
	bytes := make([]byte, 1)
	_, err := b.IoOptions.CommandReader.Read(bytes)
	if err == io.EOF {
		return nil, false
	}
	command := rune(bytes[0])

	if !b.isCommandDefined(&command) {
		return b.readCommand()
	}

	return &command, true
}

// isCommandPointerAtLast check if last command is executed or not
func (b *Brainfuck) isCommandPointerAtLast() bool {
	return b.CommandPointer == int64(len(b.Commands))
}
