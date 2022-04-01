# GoCal
A simple mathematical parser written in Go on the day before my gynaecology finals.

# Usage
Build by `go build -o gocal calc.go` and move the binary `gocal` into `/usr/local/bin`. 
`gocal -i` starts an interactive prompt, `gocal -c` takes a string as a command line argument and parses it.
Running just `gocal` waits for an input. BTW, the parser doesn't just compute the answer, it also shows the generated AST (abstract syntax tree).  
