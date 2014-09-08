
package interpreter

import (
  "github.com/Cirru/writer"
)

func (env *scope) _string(xs sequence) (ret unitype) {
  if token, ok := xs[0].(token); ok {
    ret.Type = uniString
    ret.Value = token.Text
  }
  if list, ok := xs[0].(sequence); ok {
    ret.Type = uniString
    lines := sequence{transformCode(list)}
    ret.Value = writer.MakeCode(lines)
  }
  return
}
