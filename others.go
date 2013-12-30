
package cirruGopher

import (
  "github.com/Cirru/cirru-grammar"
)

func cirruType(env *Env, xs cirru.List) (ret Object) {
  value := cirruGet(env, xs[0:1])
  if &value != nil {
    ret.Tag = "string"
    ret.Value = value.Tag
  }
  return
}
