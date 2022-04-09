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

func readInputOperation(b *Brainfuck) error {
	return nil
}

func startLoopOperation(b *Brainfuck) error {
	return nil
}

func endLoopOperation(b *Brainfuck) error {
	return nil
}
