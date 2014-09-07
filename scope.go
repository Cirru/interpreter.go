
package interpreter

import "fmt"

func (env *Env) self(xs []interface{}) (ret Object) {
  ret.Tag = cirruTable
  ret.Value = env
  return
}

func (env *Env) child(xs []interface{}) (ret Object) {
  childTable := Env{}
  childTable["parent"] = env.self(xs)
  ret.Tag = cirruTable
  ret.Value = &childTable
  // println("ret is:", ret.Value)
  return
}

func (env *Env) under(xs []interface{}) (ret Object) {
  item := env.get(xs[0:1])
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

func (env *Env) code(xs []interface{}) (ret Object) {
  ret.Tag = cirruCode
  ret.Value = &xs
  return
}

func (env *Env) eval(xs []interface{}) (ret Object) {
  switch len(xs) {
  case 1:
    if code, ok := env.get(xs[0:1]).Value.(*[]interface{}); ok {
      for _, line := range *code {
        if codeLine, ok := line.([]interface{}); ok {
          ret = Evaluate(env, codeLine)
          return
        }
      }
    }
  case 2:
    if scope, ok := env.get(xs[0:1]).Value.(*Env); ok {
      if code, ok := env.get(xs[1:2]).Value.(*[]interface{}); ok {
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