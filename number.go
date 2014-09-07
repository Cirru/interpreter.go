
package interpreter

import (
  "github.com/Cirru/parser"
  "strconv"
)

func (env *Env) _int(xs []interface{}) (ret Object) {
  if token, ok := xs[0].(parser.Token); ok {
    intNumber, err := strconv.Atoi(token.Text)
    if err != nil {
      panic(err)
    }
    ret.Tag = cirruInt
    ret.Value = intNumber
  }
  return
}

func (env *Env) _float(xs []interface{}) (ret Object) {
  if token, ok := xs[0].(parser.Token); ok {
    floatNumber, err := strconv.ParseFloat(token.Text, 64)
    if err != nil {
      panic(err)
    }
    ret.Tag = cirruFloat
    ret.Value = floatNumber
  }
  return
}
