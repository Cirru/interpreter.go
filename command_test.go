package interpreter

import (
  "testing"
  "path"
  // "io/ioutil"
)

func TestCommand(t *testing.T) {
  // files, _ := ioutil.ReadDir("cirru/")
  files := []string{"macro.cirru"}
  for _, file := range files {
    // filepath := path.Join("cirru/", file.Name())
    filepath := path.Join("cirru/", file)
    println()
    println("---> Running for:", filepath)
    println()
    err := Interpret(filepath)
    if err != nil {
      t.Errorf("Runtime error", err)
    }
  }
}
