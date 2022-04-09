package brainfuck

import "io"

const (
	data_size = 30000
)

type Brainfuck struct {
	commands       map[rune]Operation
	Data           [data_size]byte
	DataPointer    int64
	Commands       [data_size]rune
	CommandPointer int64
	IoOptions      *IoOptions
	loopStack      []Loop
}

type IoOptions struct {
	CommandReader io.Reader
	InputReader   io.Reader
	OutputWriter  io.Writer
}

type CustomOperation struct {
	Character rune
	Operation Operation
}

type Loop struct {
	Start *int64
	End   *int64
}
