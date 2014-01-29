
package cirruGopher

import (
  "github.com/Cirru/cirru-parser.go"
)

func cirruArray(env *Env, xs cirru.List) (ret Object) {
  ret.Tag = "array"
  hold := []Object{}
  for _, item := range xs {
    list := cirru.List{item}
    hold = append(hold, cirruGet(env, list))
  }
  tmp := []interface{}{}
  tmp = append(tmp, &hold)
  ret.Value = tmp[0]
  return
}
