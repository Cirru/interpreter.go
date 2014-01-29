
package cirruGopher

import (
  "github.com/Cirru/cirru-parser.go"
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
  switch len(xs) {
  case 2:
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
  case 3:
    hold := cirruGet(env, xs[0:1])
    if scope, ok := hold.Value.(*Env); ok {
      ret = cirruSet(scope, xs[1:3])
      return
    }
  default:
    stop("parameter length not correct for set")
  }
  return
}

func cirruGet(env *Env, xs cirru.List) (ret Object) {
  switch len(xs) {
  case 1:
    if token, ok := xs[0].(cirru.Token); ok {
      if value, ok := (*env)[token.Text]; ok {
        ret = value
        return
      } else {
        if parent, ok := (*env)["parent"]; ok {
          if scope, ok := parent.Value.(*Env); ok {
            ret = cirruGet(scope, xs[0:1])
            return
          }
        }
      }
      return
    }
  case 2:
    item := cirruGet(env, xs[0:1])
    if scope, ok := item.Value.(*Env); ok {
      ret = cirruGet(scope, xs[1:2])
      return
    }
  default:
    stop("length", len(xs), "is not correct")
  }
  if list, ok := xs[0].(cirru.List); ok {
    ret = Evaluate(env, list)
    return
  }
  return
}
