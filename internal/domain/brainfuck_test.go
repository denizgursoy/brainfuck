package brainfuck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrainFuck_NewBrainFuck(t *testing.T) {

	operation := func(b *Brainfuck) error {
		return nil
	}

	createNewOption := func(character rune, operation Operation) Option {
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

	t.Run("should create a brainfuck successfully", func(t *testing.T) {
		brainfuck, err := NewBrainFuck()
		assert.Nil(t, err)
		assert.NotNil(t, brainfuck)
	})

	t.Run("should have all mandatory commands", func(t *testing.T) {
		brainfuck, err := NewBrainFuck()

		assert.NotNil(t, brainfuck)
		assert.Nil(t, err)

		mandatoryOperationCharacters := []rune{'+', '-', '>', '<', '.', ',', '[', ']'}
		for _, character := range mandatoryOperationCharacters {
			assert.NotNil(t, brainfuck.commands[character])
		}

	})

	t.Run("should add users' custom operation to the brainfuck", func(t *testing.T) {
		character := '*'
		option := createNewOption(character, operation)
		brainfuck, err := NewBrainFuck(option)

		assert.NotNil(t, brainfuck.commands[character])
		assert.NotNil(t, brainfuck)
		assert.Nil(t, err)
	})

	t.Run("should not add operation with an existing character", func(t *testing.T) {
		character := '+'
		option := createNewOption(character, operation)

		brainfuck, err := NewBrainFuck(option)

		assert.Nil(t, brainfuck)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, OperationExistsError)
	})

	t.Run("should not add new operator if it is nil", func(t *testing.T) {

		character := '*'
		option := createNewOption(character, nil)

		brainfuck, err := NewBrainFuck(option)

		assert.Nil(t, brainfuck)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, OperationNilError)

	})
}
