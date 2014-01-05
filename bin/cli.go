
package main

import (
  "github.com/Cirru/cirru-gopher"
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