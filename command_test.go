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
  }
  for _, filename := range filenames {
    filepath := path.Join("cirru/", filename)
    println()
    println("---> Running for:", filepath)
    println()
    err := Interpret(filepath)
    if err != nil {
      t.Errorf("Runtime error", err)
    }
  }
}
