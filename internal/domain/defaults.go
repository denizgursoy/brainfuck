package brainfuck

func incrementOperation(b *Brainfuck) error {
	b.Data[b.DataPointer]++
	return nil
}

func decrementOperation(b *Brainfuck) error {
	b.Data[b.DataPointer]--
	return nil
}

func shiftRightOperation(b *Brainfuck) error {
	b.DataPointer++
	return nil
}

func shiftLeftOperation(b *Brainfuck) error {
	b.DataPointer--
	return nil
}

func printOperation(b *Brainfuck) error {
	bytes := make([]byte, 0)
	bytes = append(bytes, b.getCurrentCellValue())
	_, err := b.IoOptions.OutputWriter.Write(bytes)
	return err
}

func setFromUserInputOperation(b *Brainfuck) error {
	bytes := make([]byte, 1)
	_, err := b.IoOptions.InputReader.Read(bytes)
	b.Data[b.DataPointer] = bytes[0]
	return err
}

func startLoopOperation(b *Brainfuck) error {
	var loop Loop
	if len(b.loopStack) == 0 {
		loop = addNewLoop(b)
	} else {
		loop = b.loopStack[len(b.loopStack)-1]
		if *loop.Start != b.CommandPointer {
			loop = addNewLoop(b)
		}
	}

	if b.getCurrentCellValue() == 0 {
		b.CommandPointer = *loop.End
	}
	return nil
}

func addNewLoop(b *Brainfuck) Loop {
	start := b.CommandPointer
	loop := Loop{
		Start: &start,
		End:   nil,
	}
	b.loopStack = append(b.loopStack, loop)
	return loop
}

func endLoopOperation(b *Brainfuck) error {
	return nil
}
