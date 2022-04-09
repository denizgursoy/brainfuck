package brainfuck

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
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

func TestDefaults_shiftRightOperation(t *testing.T) {
	ioOptions := createIoOptions()
	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.DataPointer = 4

	err := shiftRightOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, brainfuck.DataPointer, int64(5))
}

func TestDefaults_shiftLeftOperation(t *testing.T) {
	ioOptions := createIoOptions()
	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.DataPointer = 4

	err := shiftLeftOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, brainfuck.DataPointer, int64(3))
}

func TestDefaults_printOperation(t *testing.T) {
	ioOptions := createIoOptions()

	buffer := bytes.Buffer{}
	ioOptions.OutputWriter = &buffer

	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.DataPointer = 4
	brainfuck.Data[brainfuck.DataPointer] = 16

	err := printOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, buffer.Bytes()[0], byte(16))
}

func TestDefaults_setFromUserInputOperation(t *testing.T) {
	ioOptions := createIoOptions()

	ioOptions.InputReader = strings.NewReader("input")

	brainfuck, _ := NewBrainFuck(ioOptions)

	brainfuck.DataPointer = 4
	brainfuck.Data[brainfuck.DataPointer] = 16

	err := setFromUserInputOperation(brainfuck)

	assert.Nil(t, err)
	assert.Equal(t, brainfuck.getCurrentCellValue(), byte(105))
}
