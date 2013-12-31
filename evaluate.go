
package cirruGopher

import (
  "github.com/Cirru/cirru-grammar"
)

type Object struct {
  Tag string
  Value interface{}
}

type Env map[string]Object

// Evaluate read expressions and return a single result
func Evaluate(env *Env, codeLine cirru.List) (ret Object) {
  // debugPrint(codeLine, *env)
  if len(codeLine) == 0 {
    emptyArray := Object{"list", codeLine}
    return emptyArray
  }

  head := codeLine[0]

  if token, ok := head.(cirru.Token); ok {
    // debugPrint(token.Text)
    switch token.Text {
    case "--": ret = cirruComment (env, codeLine[1:])
    case "array": ret = cirruArray (env, codeLine[1:])
    case "block": ret = cirruBlock(env, codeLine[1:])
    case "bool": ret = cirruBool(env, codeLine[1:])
    case "call": ret = cirruCall(env, codeLine[1:])
    case "child": ret = cirruChild(env, codeLine[1:])
    case "code": ret = cirruCode(env, codeLine[1:])
    case "echo": ret = cirruEcho(env, codeLine[1:])
    case "eval": ret = cirruEval(env, codeLine[1:])
    case "float": ret = cirruFloat(env, codeLine[1:])
    case "get": ret = cirruGet(env, codeLine[1:])
    case "int": ret = cirruInt(env, codeLine[1:])
    case "map": ret = cirruMap(env, codeLine[1:])
    case "print": ret = cirruPrint(env, codeLine[1:])
    case "regexp": ret = cirruRegexp(env, codeLine[1:])
    case "require": ret = cirruRequire(env, codeLine[1:])
    case "self": ret = cirruSelf(env, codeLine[1:])
    case "set": ret = cirruSet(env, codeLine[1:])
    case "string": ret = cirruString(env, codeLine[1:])
    case "type": ret = cirruType(env, codeLine[1:])
    case "under": ret = cirruUnder(env, codeLine[1:])
    default: stop(token.Text, "not found")
    }
    return
  }
  if headExpression, ok := head.(cirru.List); ok {
    debugPrint(headExpression)
    return
  }
  return
}