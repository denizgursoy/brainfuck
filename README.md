# Brainfuck


## Cell data type
I used rune as cell data type because it allows users to represent characters in Unicode

## Initial data slice size

I set initial data slice capacity to 1000. If more > operation is received, new values are appended to slice
Initial capacity should be at least one. There is also a test checking that

## Test

Test can be run with
`make test-all` or `go test -v ./...`

## Customization

New operations can be added during the creation of Brainfuck. As it can be seen in the example below,
User needs to create an option which calls `ExtendWith` method of brainfuck. `ExtendWith` method takes Operation
and the Character related to it.

```go

    ioOptions := brainfuck.IoOptions{
		CommandReader: commandFile,
		InputReader:   inputFile,
		OutputWriter:  outputFile,
	}

	multiplyOption := func(b *brainfuck.Brainfuck) error {
		return b.ExtendWith(brainfuck.CustomOperation{
			Character: '*',
			Operation: func(b *brainfuck.Brainfuck) error {
				b.Data[b.DataPointer] = b.Data[b.DataPointer] * 2
				return nil
			},
		})
	}

	brainFuck, err := brainfuck.NewBrainFuck(&ioOptions, multiplyOption)
```
