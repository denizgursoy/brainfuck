# Brainfuck


## Cell data type
I used rune as cell data type because it allows users to represent characters in Unicode

## Initial data slice size

I set initial data slice capacity to 1000. If more > operation is received, new values are appended to slice
Initial capacity should be at least one. There is also a test checking that

## Test

Test can be run with
`make test-all` or `go test -v ./...`


