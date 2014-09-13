
package interpreter

import (
  "strconv"
)

func (env *scope) float(xs sequence) (ret unitype) {
  tok, ok := xs[0].(token)
  if !ok {
    panic("float expects token")
  }
  value, err := strconv.ParseFloat(tok.Text, 64)
  if err != nil {
    panic(err)
  }
  ret.Type = uniFloat
  ret.Value = value
  return
}
