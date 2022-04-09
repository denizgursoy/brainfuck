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
		brainfuck, _ := NewBrainFuck()
		mandatoryOperationCharacters := []rune{'+', '-', '>', '<', '.', ',', '[', ']'}
		for _, character := range mandatoryOperationCharacters {
			assert.NotNil(t, brainfuck.Commands[character])
		}
	})
}
