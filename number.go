
package interpreter

import (
  "github.com/Cirru/parser"
  "strconv"
)

func (env *Env) _int(xs []interface{}) (ret unitype) {
  if token, ok := xs[0].(parser.Token); ok {
    intNumber, err := strconv.Atoi(token.Text)
    if err != nil {
      panic(err)
    }
    ret.Type = uniInt
    ret.Value = intNumber
  }
  return
}

func (env *Env) _float(xs []interface{}) (ret unitype) {
  if token, ok := xs[0].(parser.Token); ok {
    floatNumber, err := strconv.ParseFloat(token.Text, 64)
    if err != nil {
      panic(err)
    }
    ret.Type = uniFloat
    ret.Value = floatNumber
  }
  return
}
