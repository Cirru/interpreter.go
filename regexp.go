
package interpreter

import (
  "github.com/Cirru/parser"
  "regexp"
)

func cirruRegexp(env *Env, xs []interface{}) (ret Object) {
  if token, ok := xs[0].(parser.Token); ok {
    reg, err := regexp.Compile(token.Text);
    if err != nil {
      panic(err)
    }
    ret.Tag = cirruTypeRegexp
    ret.Value = reg
  }
  return
}