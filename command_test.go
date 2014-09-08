package interpreter

import (
  "testing"
  "path"
  "io/ioutil"
)

func TestCommand(t *testing.T) {
  files, _ := ioutil.ReadDir("cirru/")
  filenames := []string{}
  for _, file := range files {
    filenames = append(filenames, file.Name())
  }
  // filenames = []string{"macro.cirru"}
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
