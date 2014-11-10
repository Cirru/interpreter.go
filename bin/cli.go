
package main

import (
  "github.com/Cirru/interpreter"
  "os"
  "fmt"
)

func main() {
  if len(os.Args) < 2 {
    interpreter.Interpret("repl")
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