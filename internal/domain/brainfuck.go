package brainfuck

func NewBrainFuck(options ...Option) (*Brainfuck, error) {
	brainFuck := createNewBrainFuck()
	brainFuck.addDefaultOperations()
	if err := brainFuck.registerOptions(options); err != nil {
		return nil, err
	}
	return brainFuck, nil
}

func createNewBrainFuck() *Brainfuck {
	return &Brainfuck{
		commands: make(map[rune]Operation, 7),
	}
}

func (b *Brainfuck) addDefaultOperations() {
	_ = b.ExtendWith(CustomOperation{'+', incrementOperation})
	_ = b.ExtendWith(CustomOperation{'-', decrementOperation})
	_ = b.ExtendWith(CustomOperation{'>', shiftRightOperation})
	_ = b.ExtendWith(CustomOperation{'<', shiftLeftOperation})
	_ = b.ExtendWith(CustomOperation{'.', printOperation})
	_ = b.ExtendWith(CustomOperation{',', readInputOperation})
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
