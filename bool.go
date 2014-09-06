
package interpreter

import (
  "github.com/Cirru/parser"
)

func cirruBool(env *Env, xs []interface{}) (ret Object) {
  ret.Tag = cirruTypeBool
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
