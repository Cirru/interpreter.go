
package interpreter

import (
  "fmt"
  "github.com/Cirru/parser"
)

func (env *Env) table(xs []interface{}) (ret Object) {
  ret.Tag = cirruTable
  hold := Env{}
  for _, item := range xs {
    if pair, ok := item.([]interface{}); ok {
      name := pair[0]
      var key string
      if token, ok := name.(parser.Token); ok {
        key = token.Text
      }
      value := env.get(pair[1:2])
      hold[key] = value
    }
  }
  ret.Value = &hold
  return
}

func (env *Env) set(xs []interface{}) (ret Object) {
  switch len(xs) {
  case 2:
    value := env.get(xs[1:2])
    if token, ok := xs[0].(parser.Token); ok {
      (*env)[token.Text] = value
      return value
    }
    if list, ok := xs[0].([]interface{}); ok {
      variable := env.get(list[0:1])
      if variable.Tag == cirruString {
        if name, ok := variable.Value.(string); ok {
          (*env)[name] = value
          return value
        }
      }
    }
  case 3:
    hold := env.get(xs[0:1])
    if scope, ok := hold.Value.(*Env); ok {
      ret = scope.set(xs[1:3])
      return
    }
  default:
    panic("parameter length not correct for set")
  }
  return
}

func (env *Env) get(xs []interface{}) (ret Object) {
  switch len(xs) {
  case 1:
    if token, ok := xs[0].(parser.Token); ok {
      if value, ok := (*env)[token.Text]; ok {
        ret = value
        return
      } else {
        if parent, ok := (*env)["parent"]; ok {
          if scope, ok := parent.Value.(*Env); ok {
            ret = scope.get(xs[0:1])
            return
          }
        }
      }
      return
    }
  case 2:
    item := env.get(xs[0:1])
    if scope, ok := item.Value.(*Env); ok {
      ret = scope.get(xs[1:2])
      return
    }
  default:
    panic(fmt.Sprintf("length %s is not correct", len(xs)))
  }
  if list, ok := xs[0].([]interface{}); ok {
    ret = Evaluate(env, list)
    return
  }
  return
}
