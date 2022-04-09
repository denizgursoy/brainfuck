package brainfuck

func NewBrainFuck(options ...Options) (*Brainfuck, error) {
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
	b.commands['+'] = incrementOperation
	b.commands['-'] = decrementOperation
	b.commands['>'] = shiftRightOperation
	b.commands['<'] = shiftLeftOperation
	b.commands['.'] = printOperation
	b.commands[','] = readInputOperation
	b.commands['['] = startLoopOperation
	b.commands[']'] = endLoopOperation
}
func (b *Brainfuck) registerOptions(options []Options) error {
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

func incrementOperation(b *Brainfuck) error {
	return nil
}

func decrementOperation(b *Brainfuck) error {
	return nil
}

func shiftRightOperation(b *Brainfuck) error {
	return nil
}

func shiftLeftOperation(b *Brainfuck) error {
	return nil
}

func printOperation(b *Brainfuck) error {
	return nil
}

func readInputOperation(b *Brainfuck) error {
	return nil
}

func startLoopOperation(b *Brainfuck) error {
	return nil
}

func endLoopOperation(b *Brainfuck) error {
	return nil
}
