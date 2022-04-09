package brainfuck

import (
	"io"
)

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
		Data:       [30000]byte{},
		Commands:   make([]rune, 0),
		IoOptions:  io,
		loopStack:  make([]*Loop, 0),
	}
	return &brainfuck, nil
}

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

func (b *Brainfuck) registerOptions(options []Option) error {
	for _, customOperation := range options {
		if err := customOperation(b); err != nil {
			return err
		}
	}
	return nil
}

func (b *Brainfuck) ExtendWith(operation CustomOperation) error {

	if operation.Operation == nil {
		return OperationNilError
	}

	if b.operations[operation.Character] != nil {
		return OperationExistsError
	}

	b.operations[operation.Character] = operation.Operation
	return nil
}

func (b *Brainfuck) getCurrentCellValue() byte {
	return b.Data[b.DataPointer]
}

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

func (b *Brainfuck) performOperation(command *rune) error {
	operation := b.operations[*command]
	return operation(b)
}

//func (b *Brainfuck) isCommandDefined(c *rune) bool {
//
//	for _, v := range b.Commands {
//		if v == *c {
//			return true
//		}
//	}
//
//	return false
//}

func (b *Brainfuck) addNewCommand(command *rune) {
	b.Commands = append(b.Commands, *command)
}

func (b *Brainfuck) readCommand() (*rune, bool) {
	bytes := make([]byte, 1)
	_, err := b.IoOptions.CommandReader.Read(bytes)
	if err == io.EOF {
		return nil, false
	}
	command := rune(bytes[0])
	return &command, true
}

func (b *Brainfuck) isCommandPointerAtLast() bool {
	return b.CommandPointer == int64(len(b.Commands))
}
