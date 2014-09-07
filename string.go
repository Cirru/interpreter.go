
package interpreter

import (
  "github.com/Cirru/parser"
  "github.com/Cirru/writer"
)

func (env *Env) _string(xs []interface{}) (ret unitype) {
  if token, ok := xs[0].(parser.Token); ok {
    ret.Type = uniString
    ret.Value = token.Text
  }
  if list, ok := xs[0].([]interface{}); ok {
    ret.Type = uniString
    lines := []interface{}{transformCode(list)}
    ret.Value = writer.MakeCode(lines)
  }
  return
}

func makeUniString(str string) unitype {
  return unitype{uniString, str}
}

func makeUniInt(n int) unitype {
  return unitype{uniInt, n}
}