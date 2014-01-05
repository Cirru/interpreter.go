
package cirruGopher

import (
  "github.com/Cirru/cirru-grammar"
)

func cirruString(env *Env, xs cirru.List) (ret Object) {
  if token, ok := xs[0].(cirru.Token); ok {
    ret.Tag = "string"
    ret.Value = token.Text
  }
  if list, ok := xs[0].(cirru.List); ok {
    ret.Tag = "string"
    ret.Value = codeString(list, 0)
  }
  return
}
