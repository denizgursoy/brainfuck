package brainfuck

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

var (
	operation = func(b *Brainfuck) error {
		return nil
	}

	createNewOption = func(character rune, operation Operation) Option {
		customOperation := CustomOperation{
			Character: character,
			Operation: operation,
		}
		option := func(b *Brainfuck) error {
			if err := b.ExtendWith(customOperation); err != nil {
				return err
			}
			return nil
		}

		return option
	}
	createIoOptions = func() *IoOptions {
		buffer := bytes.Buffer{}
		return &IoOptions{
			CommandReader: strings.NewReader(""),
			InputReader:   strings.NewReader(""),
			OutputWriter:  &buffer,
		}
	}
)

func TestBrainFuck_NewBrainFuck(t *testing.T) {

	t.Run("should create a brainfuck successfully", func(t *testing.T) {
		brainfuck, err := NewBrainFuck(createIoOptions())
		assert.Nil(t, err)
		assert.NotNil(t, brainfuck)
	})

	t.Run("should have all mandatory operations", func(t *testing.T) {
		brainfuck, err := NewBrainFuck(createIoOptions())

		assert.NotNil(t, brainfuck)
		assert.Nil(t, err)

		mandatoryOperationCharacters := []rune{'+', '-', '>', '<', '.', ',', '[', ']'}
		for _, character := range mandatoryOperationCharacters {
			assert.NotNil(t, brainfuck.operations[character])
		}

	})

	t.Run("should add users' custom operation to the brainfuck", func(t *testing.T) {
		character := '*'
		option := createNewOption(character, operation)
		brainfuck, err := NewBrainFuck(createIoOptions(), option)

		assert.NotNil(t, brainfuck.operations[character])
		assert.NotNil(t, brainfuck)
		assert.Nil(t, err)
	})

	t.Run("should not add operation with an existing character", func(t *testing.T) {
		character := '+'
		option := createNewOption(character, operation)

		brainfuck, err := NewBrainFuck(createIoOptions(), option)

		assert.Nil(t, brainfuck)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, OperationExistsError)
	})

	t.Run("should not add new operator if it is nil", func(t *testing.T) {

		character := '*'
		option := createNewOption(character, nil)

		brainfuck, err := NewBrainFuck(createIoOptions(), option)

		assert.Nil(t, brainfuck)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, OperationNilError)

	})

	t.Run("should have all io options", func(t *testing.T) {
		t.Run("should have a command reader", func(t *testing.T) {
			ioOptions := createIoOptions()
			ioOptions.CommandReader = nil

			brainfuck, err := NewBrainFuck(ioOptions)

			assert.NotNil(t, err)
			assert.ErrorIs(t, err, CommandReaderNilError)
			assert.Nil(t, brainfuck)
		})

		t.Run("should have a command reader", func(t *testing.T) {
			ioOptions := createIoOptions()
			ioOptions.InputReader = nil

			brainfuck, err := NewBrainFuck(ioOptions)

			assert.NotNil(t, err)
			assert.ErrorIs(t, err, InputReaderNilError)
			assert.Nil(t, brainfuck)
		})

		t.Run("should have a output writer", func(t *testing.T) {
			ioOptions := createIoOptions()
			ioOptions.OutputWriter = nil

			brainfuck, err := NewBrainFuck(ioOptions)

			assert.NotNil(t, err)
			assert.ErrorIs(t, err, OutputWriterNilError)
			assert.Nil(t, brainfuck)
		})

		t.Run("should initialize all slice in the brainfuck", func(t *testing.T) {
			ioOptions := createIoOptions()

			brainfuck, err := NewBrainFuck(ioOptions)
			assert.Nil(t, err)
			assert.NotNil(t, brainfuck.operations)
			assert.NotNil(t, brainfuck.Commands)

			assert.NotNil(t, brainfuck.Data)
			assert.Len(t, brainfuck.Data, InitialCapacity)
		})
	})
}

func TestBrainFuck_calculateCommandToExecute(t *testing.T) {
	t.Run("should read from command reader if command pointer is at last", func(t *testing.T) {
		ioOptions := createIoOptions()
		ioOptions.CommandReader = strings.NewReader(">")
		brainfuck, err := NewBrainFuck(ioOptions)
		brainfuck.getCommandToExecute()

		assert.Nil(t, err)
		assert.Equal(t, brainfuck.Commands[0], '>')
	})
}

func TestBrainFuck_Start(t *testing.T) {
	helloWorld, _ := ioutil.ReadFile("bf/Hello.bf")
	sixHundred, _ := ioutil.ReadFile("bf/666.bf")

	table := []struct {
		Input  string
		Output string
	}{
		{
			Input:  string(helloWorld),
			Output: "Hello World!\n",
		},
		{
			Input:  string(sixHundred),
			Output: "666\n",
		},
	}

	for _, testCase := range table {

		ioOptions := createIoOptions()
		buffer := bytes.Buffer{}

		ioOptions.CommandReader = strings.NewReader(testCase.Input)
		ioOptions.OutputWriter = &buffer

		brainfuck, _ := NewBrainFuck(ioOptions)
		_ = brainfuck.Start()

		assert.Equal(t, testCase.Output, string(buffer.Bytes()))
	}
}
