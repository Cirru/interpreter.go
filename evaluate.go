
package interpreter

import (
  "github.com/Cirru/parser"
  "fmt"
)

type Object struct {
  Tag cirruName
  Value interface{}
}

type Env map[string]Object

// Evaluate read expressions and return a single result
func Evaluate(env *Env, xs []interface{}) (ret Object) {
  if len(xs) == 0 {
    emptyArray := Object{cirruArray, xs}
    return emptyArray
  }

  if token, ok := xs[0].(parser.Token); ok {
    switch token.Text {
    case "--": ret = env.comment(xs[1:])
    case "array": ret = env.array(xs[1:])
    case "function":  ret = env.function(xs[1:])
    case "bool":      ret = env._bool(xs[1:])
    case "call":      ret = env.call(xs[1:])
    case "child":     ret = env.child(xs[1:])
    case "code":      ret = env.code(xs[1:])
    case "eval":      ret = env.eval(xs[1:])
    case "float":     ret = env._float(xs[1:])
    case "get":       ret = env.get(xs[1:])
    case "int":       ret = env._int(xs[1:])
    case "table":     ret = env.table(xs[1:])
    case "print":     ret = env._print(xs[1:])
    case "regexp":    ret = env.regexp(xs[1:])
    case "require":   ret = env.require(xs[1:])
    case "self":      ret = env.self(xs[1:])
    case "set":       ret = env.set(xs[1:])
    case "string":    ret = env._string(xs[1:])
    case "type":      ret = env._type(xs[1:])
    case "under":     ret = env.under(xs[1:])
    default:
      ret = userCall(env, xs)
    }
    return
  }
  if headExpression, ok := xs[0].([]interface{}); ok {
    fmt.Println(headExpression)
    return
  }
  return
}

func userCall(env *Env, xs []interface{}) (ret Object) {
  head := env.get(xs[0:1])
  if head.Tag == cirruFunction {
    ret = env.call(xs)
  }
  return
}