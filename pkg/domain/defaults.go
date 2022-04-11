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
	if b.DataPointer == int64(len(b.Data))-1 {
		b.Data = append(b.Data, 0)
	}
	b.DataPointer++
	return nil
}

func shiftLeftOperation(b *Brainfuck) error {
	if b.DataPointer == 0 {
		return ShiftLeftNoSpaceError
	}
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
	var loop *Loop
	if len(b.loopStack) == 0 {
		loop = addNewLoop(b)
	} else {
		loop = b.loopStack[len(b.loopStack)-1]
		if *loop.Start != b.CommandPointer {
			loop = addNewLoop(b)
		}
	}

	if b.getCurrentCellValue() == 0 {
		if loop.End == nil {
			return LoopEndIsNotFoundError
		}
		b.CommandPointer = *loop.End
	}
	return nil
}

func endLoopOperation(b *Brainfuck) error {

	if len(b.loopStack) == 0 {
		return LoopEndInvalidError
	}

	if b.getCurrentCellValue() > 0 {
		loop := peekStack(b)
		end := b.CommandPointer
		loop.End = &end
		b.CommandPointer = *loop.Start
	} else {
		popStack(b)
	}
	return nil
}

func addNewLoop(b *Brainfuck) *Loop {
	start := b.CommandPointer
	loop := &Loop{
		Start: &start,
		End:   nil,
	}
	b.loopStack = append(b.loopStack, loop)
	return loop
}

func peekStack(b *Brainfuck) *Loop {
	return b.loopStack[len(b.loopStack)-1]
}

func popStack(b *Brainfuck) {
	b.loopStack = b.loopStack[:len(b.loopStack)-1]
}
