
// Package cirruGopher is a small interpreter of parser.
// It is based on `parser`.
package interpreter

import (
  "io/ioutil"
  "github.com/Cirru/parser"
)

// Interpret takes result from `parser.Parse` and run in context.
func Interpret(filepath string) error {
  moduleCenter = Env{}
  scope := Env{}
  exports := Env{}
  scope[makeUniString("filepath")] = generateString(filepath)
  ret := generateTable(&exports)
  scope[makeUniString("exports")] = ret
  moduleCenter[makeUniString(filepath)] = ret

  codeByte, err := ioutil.ReadFile(filepath)
  if err != nil {
    panic(err)
  }

  p := parser.NewParser()
  p.Filename(filepath)
  for _, c := range codeByte {
    p.Read(rune(c))
  }
  p.Complete()

  ast := p.ToArray()

  for _, line := range ast {
    if list, ok := line.([]interface{}); ok {
      Evaluate(&scope, list)
    }
  }
  return nil
}
