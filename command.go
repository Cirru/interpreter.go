
// Package cirruGopher is a small interpreter of Cirru.
// It is based on `cirru-parser.go`.
package cirruGopher

import (
  "io/ioutil"
  "github.com/Cirru/cirru-parser.go"
)

// Interpret takes result from `cirru.Parse` and run in context.
func Interpret(filepath string) error {
  moduleCenter = Env{}
  scope := Env{}
  exports := Env{}
  scope["filepath"] = generateString(filepath)
  ret := generateMap(&exports)
  scope["exports"] = ret
  moduleCenter[filepath] = ret
  
  codeByte, err := ioutil.ReadFile(filepath)
  if err != nil {
    panic(err)
  }
  code := string(codeByte)
  ast := cirru.Parse(code, filepath)

  for _, line := range ast {
    if list, ok := line.(cirru.List); ok {
      Evaluate(&scope, list)
    }
  }
  return nil
}
