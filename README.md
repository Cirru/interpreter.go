
Cirru Interpreter in Go
------

Run Cirru in Go.

See [master/code/](https://github.com/Cirru/interpreter/tree/master/code) for demos.

Here's how you can use it:

```go
package main

import (
  "github.com/Cirru/interpreter"
  "os"
  "fmt"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Please specify filename")
    os.Exit(1)
  }
  first := os.Args[1]
  if _, err := os.Stat(first); err == nil {
    interpreter.Interpret(first)
  } else {
    fmt.Println("Please specify a file")
    os.Exit(1)
  }
}
```

And by running `go build bin/cli.go` you may get a binary that runs Cirru code.

[![GoDoc](https://godoc.org/github.com/Cirru/interpreter?status.png)](https://godoc.org/github.com/Cirru/interpreter)

### License

MIT