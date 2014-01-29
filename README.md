
Cirru Interpreter in Go
------

Run Cirru in Go. A rewrite of [`cirru-interpreter`][interpreter] in Go.

See [master/code/](https://github.com/Cirru/cirru-interpreter.go/tree/master/code) for demos.

Here's how you can use it:

```go
package main

import (
  "github.com/Cirru/cirru-interpreter.go"
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
    cirruGopher.Interpret(first)
  } else {
    fmt.Println("Please specify a file")
    os.Exit(1)
  }
}
```

And by running `go build bin/cli.go` you may get a binary that runs Cirru code.

[interpreter]: https://github.com/Cirru/cirru-interpreter.coffee

[![GoDoc](https://godoc.org/github.com/Cirru/cirru-interpreter.go?status.png)](https://godoc.org/github.com/Cirru/cirru-interpreter.go)

### License

MIT