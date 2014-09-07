
package interpreter

import (
  "github.com/Cirru/parser"
  "fmt"
)

type Object struct {
  Tag cirruTypeName
  Value interface{}
}

type Env map[string]Object

// Evaluate read expressions and return a single result
func Evaluate(env *Env, xs []interface{}) (ret Object) {
  if len(xs) == 0 {
    emptyArray := Object{cirruTypeArray, xs}
    return emptyArray
  }

  if token, ok := xs[0].(parser.Token); ok {
    switch token.Text {
    case "--": ret = cirruComment (env, xs[1:])
    case "array": ret = cirruArray (env, xs[1:])
    case "function": ret = cirruFunction(env, xs[1:])
    case "bool": ret = cirruBool(env, xs[1:])
    case "call": ret = cirruCall(env, xs[1:])
    case "child": ret = cirruChild(env, xs[1:])
    case "code": ret = cirruCode(env, xs[1:])
    case "eval": ret = cirruEval(env, xs[1:])
    case "float": ret = cirruFloat(env, xs[1:])
    case "get": ret = cirruGet(env, xs[1:])
    case "int": ret = cirruInt(env, xs[1:])
    case "table": ret = cirruTable(env, xs[1:])
    case "print": ret = cirruPrint(env, xs[1:])
    case "regexp": ret = cirruRegexp(env, xs[1:])
    case "require": ret = cirruRequire(env, xs[1:])
    case "self": ret = cirruSelf(env, xs[1:])
    case "set": ret = cirruSet(env, xs[1:])
    case "string": ret = cirruString(env, xs[1:])
    case "type": ret = cirruType(env, xs[1:])
    case "under": ret = cirruUnder(env, xs[1:])
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
  head := cirruGet(env, xs[0:1])
  if head.Tag == cirruTypeFunction {
    ret = cirruCall(env, xs)
  }
  return
}