package interpreter

import (
  "testing"
  "path"
)

func TestCommand(t *testing.T) {
  filenames := []string{
    "array.cirru",
    "bool.cirru",
    "number.cirru",
    "regexp.cirru",
    "string.cirru",
    "table.cirru",
    "stdio.cirru",
    "function.cirru",
    "ctrl.cirru",
    "require.cirru",
    "method.cirru",
    "closure.cirru",
    "math.cirru",
  }
  for _, filename := range filenames {
    filepath := path.Join("cirru/", filename)
    println()
    println("---- Interpreting:", filepath)
    println()
    Interpret(filepath)
  }
}
