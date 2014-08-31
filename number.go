
package interpreter

import (
  "github.com/Cirru/parser"
  "strconv"
)

func cirruInt(env *Env, xs []interface{}) (ret Object) {
  if token, ok := xs[0].(parser.Token); ok {
    intNumber, err := strconv.Atoi(token.Text)
    if err != nil {
      panic(err)
    }
    ret.Tag = "int"
    ret.Value = intNumber
  }
  return
}

func cirruFloat(env *Env, xs []interface{}) (ret Object) {
  if token, ok := xs[0].(parser.Token); ok {
    floatNumber, err := strconv.ParseFloat(token.Text, 64)
    if err != nil {
      panic(err)
    }
    ret.Tag = "float"
    ret.Value = floatNumber
  }
  return
}
