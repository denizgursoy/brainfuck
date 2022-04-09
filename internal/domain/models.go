package brainfuck

type Brainfuck struct {
	commands map[rune]Operation
}
