
package interpreter

import (
  "strconv"
)

func (env *scope) _int(xs sequence) (ret unitype) {
  if t, ok := xs[0].(token); ok {
    intNumber, err := strconv.Atoi(t.Text)
    if err != nil {
      panic(err)
    }
    ret.Type = uniInt
    ret.Value = intNumber
  } else {
    panic("can not parse as number")
  }
  return
}

func (env *scope) _float(xs sequence) (ret unitype) {
  if t, ok := xs[0].(token); ok {
    floatNumber, err := strconv.ParseFloat(t.Text, 64)
    if err != nil {
      panic(err)
    }
    ret.Type = uniFloat
    ret.Value = floatNumber
  }
  return
}
