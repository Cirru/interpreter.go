
package interpreter

import (
  "github.com/Cirru/parser"
)

func (env *Env) _bool(xs []interface{}) (ret unitype) {
  ret.Type = uniBool
  ret.Value = false
  if token, ok := xs[0].(parser.Token); ok {
    trueValues := []string{"true", "yes", "right", "1"}
    for _, text := range trueValues {
      if text == token.Text {
        ret.Value = true
      }
    }
    return
  }
  return
}
