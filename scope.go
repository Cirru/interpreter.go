
package interpreter

import "fmt"

func cirruSelf(env *Env, xs []interface{}) (ret Object) {
  ret.Tag = cirruTypeMap
  ret.Value = env
  return
}

func cirruChild(env *Env, xs []interface{}) (ret Object) {
  childMap := Env{}
  childMap["parent"] = cirruSelf(env, xs)
  ret.Tag = cirruTypeMap
  ret.Value = &childMap
  // println("ret is:", ret.Value)
  return
}

func cirruUnder(env *Env, xs []interface{}) (ret Object) {
  item := cirruGet(env, xs[0:1])
  if scope, ok := item.Value.(*Env); ok {
    for _, exp := range xs[1:] {
      if list, ok := exp.([]interface{}); ok {
        ret = Evaluate(scope, list)
      }
    }
  } else {
    fmt.Println("no scope", item.Value, xs[1:])
  }
  return
}

func cirruCode(env *Env, xs []interface{}) (ret Object) {
  ret.Tag = cirruTypeCode
  ret.Value = &xs
  return
}

func cirruEval(env *Env, xs []interface{}) (ret Object) {
  switch len(xs) {
  case 1:
    if code, ok := cirruGet(env, xs[0:1]).Value.(*[]interface{}); ok {
      for _, line := range *code {
        if codeLine, ok := line.([]interface{}); ok {
          ret = Evaluate(env, codeLine)
          return
        }
      }
    }
  case 2:
    if scope, ok := cirruGet(env, xs[0:1]).Value.(*Env); ok {
      if code, ok := cirruGet(env, xs[1:2]).Value.(*[]interface{}); ok {
        for _, line := range *code {
          if codeLine, ok := line.([]interface{}); ok {
            ret = Evaluate(scope, codeLine)
          }
        }
      }
    }
  }
  return
}