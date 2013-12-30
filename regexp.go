
package cirruGopher

import (
  "github.com/Cirru/cirru-grammar"
  "regexp"
)

func cirruRegexp(env *Env, xs cirru.List) (ret Object) {
  if token, ok := xs[0].(cirru.Token); ok {
    reg, err := regexp.Compile(token.Text);
    if err != nil {
      panic(err)
    }
    ret.Tag = "regexp"
    ret.Value = reg
  }
  return
}