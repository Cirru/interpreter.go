
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
)

func cirruSelf(env *Env, xs cirru.List) (ret Object) {
  ret.Tag = "map"
  ret.Value = env
  return
}

func cirruChild(env *Env, xs cirru.List) (ret Object) {
  childMap := Env{}
  childMap["parent"] = cirruSelf(env, xs)
  ret.Tag = "map"
  ret.Value = &childMap
  // println("ret is:", ret.Value)
  return
}

func cirruUnder(env *Env, xs cirru.List) (ret Object) {
  item := cirruGet(env, xs[0:1])
  // debugPrint("item is:", item.Value)
  if scope, ok := item.Value.(*Env); ok {
    // debugPrint("scope is:", xs[1])
    for _, exp := range xs[1:] {
      if list, ok := exp.(cirru.List); ok {
        ret = Evaluate(scope, list)
      }
    }
  } else {
    debugPrint("no scope", item.Value, xs[1:])
  }
  return
}
