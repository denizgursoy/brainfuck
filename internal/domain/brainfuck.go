package brainfuck

func NewBrainFuck() (*Brainfuck, error) {
	brainFuck := createNewBrainFuck()
	brainFuck.addDefaultOperations()
	return brainFuck, nil
}

func createNewBrainFuck() *Brainfuck {
	return &Brainfuck{
		Commands: make(map[rune]Operation, 7),
	}
}

func (b *Brainfuck) addDefaultOperations() {
	b.Commands['+'] = incrementOperation
	b.Commands['-'] = decrementOperation
	b.Commands['>'] = shiftRightOperation
	b.Commands['<'] = shiftLeftOperation
	b.Commands['.'] = printOperation
	b.Commands[','] = readInputOperation
	b.Commands['['] = startLoopOperation
	b.Commands[']'] = endLoopOperation
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
