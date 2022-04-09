package brainfuck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBrainFuck_NewBrainFuck(t *testing.T) {

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
		operationCharacter := '*'
		operation := CustomOperation{
			Character: operationCharacter,
			Operation: func(b *Brainfuck) error {
				return nil
			},
		}
		option := func(b *Brainfuck) error {
			if err := b.ExtendWith(operation); err != nil {
				return err
			}
			return nil
		}
		brainfuck, err := NewBrainFuck(option)

		assert.NotNil(t, brainfuck.commands[operationCharacter])
		assert.NotNil(t, brainfuck)
		assert.Nil(t, err)
	})

	t.Run("should not add operation with an existing character", func(t *testing.T) {
		operationCharacter := '+'
		operation := CustomOperation{
			Character: operationCharacter,
			Operation: func(b *Brainfuck) error {
				return nil
			},
		}
		option := func(b *Brainfuck) error {
			if err := b.ExtendWith(operation); err != nil {
				return err
			}
			return nil
		}
		brainfuck, err := NewBrainFuck(option)

		assert.Nil(t, brainfuck)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, OperationExistsError)
	})
}
