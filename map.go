
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
)

func cirruMap(env *Env, xs cirru.List) (ret Object) {
  ret.Tag = "map"
  hold := Env{}
  for _, item := range xs {
    if pair, ok := item.(cirru.List); ok {
      name := pair[0]
      var key string
      if token, ok := name.(cirru.Token); ok {
        key = token.Text
      }
      value := cirruGet(env, pair[1:2])
      hold[key] = value
    }
  }
  ret.Value = &hold
  return
}

func cirruSet(env *Env, xs cirru.List) (ret Object) {
  // debugPrint("cirruSet:", xs)
  value := cirruGet(env, xs[1:2])
  if token, ok := xs[0].(cirru.Token); ok {
    (*env)[token.Text] = value
    return value
  }
  if list, ok := xs[0].(cirru.List); ok {
    variable := cirruGet(env, list[0:1])
    if variable.Tag == "string" {
      if name, ok := variable.Value.(string); ok {
        (*env)[name] = value
        return value
      }
    }
  }
  return
}

func cirruGet(env *Env, xs cirru.List) (ret Object) {
  if token, ok := xs[0].(cirru.Token); ok {
    if value, ok := (*env)[token.Text]; ok {
      ret = value
      return
    } else {
      if parent, ok := (*env)["parent"]; ok {
        if parent.Tag == "map" {
          if scope, ok := parent.Value.(*Env); ok {
            ret = cirruGet(scope, xs[0:1])
            return
          }
        }
      }
    }
    return
  }
  if list, ok := xs[0].(cirru.List); ok {
    ret = Evaluate(env, list)
    return
  }
  return
}
