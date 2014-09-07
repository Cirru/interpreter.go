
package interpreter

import (
  "github.com/Cirru/parser"
  "github.com/Cirru/writer"
)

func (env *Env) _string(xs []interface{}) (ret Object) {
  if token, ok := xs[0].(parser.Token); ok {
    ret.Tag = cirruString
    ret.Value = token.Text
  }
  if list, ok := xs[0].([]interface{}); ok {
    ret.Tag = cirruString
    lines := []interface{}{transformCode(list)}
    ret.Value = writer.MakeCode(lines)
  }
  return
}
