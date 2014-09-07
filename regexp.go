
package interpreter

import (
  "github.com/Cirru/parser"
  "regexp"
)

func (env *Env) regexp(xs []interface{}) (ret unitype) {
  if token, ok := xs[0].(parser.Token); ok {
    reg, err := regexp.Compile(token.Text);
    if err != nil {
      panic(err)
    }
    ret.Type = cirruRegexp
    ret.Value = reg
  }
  return
}