package brainfuck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaults_incrementOperation(t *testing.T) {

	ioOptions := createIoOptions()
	brainfuck, _ := NewBrainFuck(ioOptions)

	err := incrementOperation(brainfuck)
	assert.Nil(t, err)
	assert.Equal(t, brainfuck.getCurrentCellValue(), byte(1))
}

func TestDefaults_decrementOperation(t *testing.T) {
	ioOptions := createIoOptions()
	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.Data[brainfuck.DataPointer] = 5

	err := decrementOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, brainfuck.getCurrentCellValue(), byte(4))
}
