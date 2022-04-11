package brainfuck

import "io"

// Brainfuck is the main struct that holds all the data needed to run operations successfully,
// Data holds cell values. DataPointer show the current cell location which is being modified.
// Commands slice stores all the command in order they are added.
// CommandPointer show the last command executed.
// loopStack holds all the loops and nested loops which are currently executed.
// IoOptions stores all readers and writers needed.
type Brainfuck struct {
	operations     map[rune]Operation
	Data           []rune
	DataPointer    int64
	Commands       []rune
	CommandPointer int64
	IoOptions      *IoOptions
	loopStack      []*Loop
}

// IoOptions holds the all readers or writers that a brainfuck needs.
// CommandReader reads new command, InputReader reads user input with ',' command,
// OutputWriter writes current cell value with '.' command.
type IoOptions struct {
	CommandReader io.Reader
	InputReader   io.Reader
	OutputWriter  io.Writer
}

// CustomOperation holds a Character that will be read from CommandReader.
// Operation is the function that will be executed is the Character is read from CommandReader.
type CustomOperation struct {
	Character rune
	Operation Operation
}

// Loop holds data of a loop created by the commands.
// Start stores the  position '[' command on Commands slice.
// End stores the  position ']' command on Commands slice.
type Loop struct {
	Start *int64
	End   *int64
}
