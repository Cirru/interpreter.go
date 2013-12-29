
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
)

func cirruBool(env *Env, xs cirru.List) (ret Object) {
  ret.Tag = "bool"
  ret.Value = false
  if token, ok := xs[0].(cirru.Token); ok {
    trueValues := []string{"true", "yes", "riight", "1"}
    for _, text := range trueValues {
      if text == token.Text {
        ret.Value = true
      }
    }
    return
  }
  return
}
