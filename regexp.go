
package interpreter

import (
  "regexp"
)

func (env *scope) regexp(xs sequence) (ret unitype) {
  if t, ok := xs[0].(token); ok {
    reg, err := regexp.Compile(t.Text);
    if err != nil {
      panic(err)
    }
    ret.Type = uniRegexp
    ret.Value = reg
  }
  return
}