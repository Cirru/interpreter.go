package cirruGopher

import (
  "testing"
  "path"
  "io/ioutil"
)

func TestCommand(t *testing.T) {
  files, _ := ioutil.ReadDir("cirru/")
  for _, file := range files {
    filepath := path.Join("cirru/", file.Name())
    println()
    println("---> Running for:", filepath)
    println()
    err := Interpret(filepath)
    if err != nil {
      t.Errorf("Runtime error", err)
    }
  }
}
