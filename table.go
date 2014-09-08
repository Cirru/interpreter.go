
package interpreter

import (
  "fmt"
)

func (env *scope) table(xs sequence) (ret unitype) {
  ret.Type = uniTable
  hold := scope{}
  for _, item := range xs {
    if pair, ok := item.(sequence); ok {
      name := pair[0]
      var key string
      if token, ok := name.(token); ok {
        key = token.Text
      }
      value := env.get(pair[1:2])
      hold[uni(key)] = value
    }
  }
  ret.Value = &hold
  return
}

func (env *scope) set(xs sequence) (ret unitype) {
  switch len(xs) {
  case 2:
    value := env.get(xs[1:2])
    if token, ok := xs[0].(token); ok {
      (*env)[uni(token.Text)] = value
      return value
    }
    if list, ok := xs[0].(sequence); ok {
      variable := env.get(list[0:1])
      if variable.Type == uniString {
        if name, ok := variable.Value.(string); ok {
          (*env)[uni(name)] = value
          return value
        }
      }
    }
  case 3:
    hold := env.get(xs[0:1])
    if scope, ok := hold.Value.(*scope); ok {
      ret = scope.set(xs[1:3])
      return
    }
  default:
    panic("parameter length not correct for set")
  }
  return
}

func (env *scope) get(xs sequence) (ret unitype) {
  switch len(xs) {
  case 1:
    if token, ok := xs[0].(token); ok {
      if value, ok := (*env)[uni(token.Text)]; ok {
        ret = value
        return
      } else {
        if parent, ok := (*env)[uni("parent")]; ok {
          if scope, ok := parent.Value.(*scope); ok {
            ret = scope.get(xs[0:1])
            return
          }
        }
      }
      return unitype{uniNil, nil}
    }
  case 2:
    item := env.get(xs[0:1])
    if scope, ok := item.Value.(*scope); ok {
      ret = scope.get(xs[1:2])
      return
    }
  default:
    panic(fmt.Sprintf("length %s is not correct", len(xs)))
  }
  if list, ok := xs[0].(sequence); ok {
    ret = Evaluate(env, list)
    return
  }
  return
}
