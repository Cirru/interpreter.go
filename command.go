
// Package cirruGopher is a small interpreter of Cirru.
// It is based on `cirru-grammar`.
package cirruGopher

import (
  "io/ioutil"
  "github.com/Cirru/cirru-grammar"
)

// Interpret takes result from `cirru.Parse` and run in context.
func Interpret() error {
  filename := "code/block.cr"
  codeByte, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }
  code := string(codeByte)
  ast := cirru.Parse(code, filename)
  globalEnv := Env{}
  for _, line := range ast {
    if list, ok := line.(cirru.List); ok {
      Evaluate(&globalEnv, list)
    }
  }
  return nil
}
