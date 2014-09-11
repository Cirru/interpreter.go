
package interpreter

import (
  "regexp"
)

func (env *scope) regexp(xs sequence) (ret unitype) {
  tok, ok := xs[0].(token)
  if !ok {
    panic("regexp excepts token")
  }
  reg, err := regexp.Compile(tok.Text)
  if err != nil {
    panic(err)
  }
  ret.Type = uniRegexp
  ret.Value = reg
  return
}