package brainfuck

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
		commands:  make(map[rune]Operation, 8),
		Data:      [30000]byte{},
		Commands:  [30000]rune{},
		IoOptions: io,
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

	if b.commands[operation.Character] != nil {
		return OperationExistsError
	}

	b.commands[operation.Character] = operation.Operation
	return nil
}

func (b *Brainfuck) getCurrentCellValue() byte {
	return b.Data[b.DataPointer]
}
