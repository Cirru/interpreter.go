
package interpreter

import (
  "github.com/Cirru/parser"
)

func cirruString(env *Env, xs []interface{}) (ret Object) {
  if token, ok := xs[0].(parser.Token); ok {
    ret.Tag = "string"
    ret.Value = token.Text
  }
  if list, ok := xs[0].([]interface{}); ok {
    ret.Tag = "string"
    ret.Value = codeString(list, 0)
  }
  return
}
