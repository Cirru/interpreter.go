
// Package cirruGopher is a small interpreter of parser.
// It is based on `parser`.
package interpreter

import (
  "io/ioutil"
  "github.com/Cirru/parser"
)

// Interpret takes result from `parser.Parse` and run in context.
func Interpret(filepath string) error {
  moduleCenter = scope{}
  fileScope := scope{}
  exports := scope{}
  fileScope[uni("filepath")] = uni(filepath)
  ret := uni(&exports)
  fileScope[uni("exports")] = ret
  moduleCenter[uni(filepath)] = ret

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

  ast := toSequence(p.ToArray())

  for _, line := range ast {
    seq, _ := line.(sequence)
    Evaluate(&fileScope, seq)
  }
  return nil
}

func toSequence(xs []interface{}) (ret sequence) {
  for _, child := range xs {
    if seq, ok := child.([]interface{}); ok {
      ret = append(ret, toSequence(seq))
    } else if t, ok := child.(parser.Token); ok {
      ret = append(ret, token(t))
    } else {
      panic("got unknown type from code")
    }
  }
  return
}