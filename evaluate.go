
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
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
    case "echo":      ret = cirruEcho    (env, codeLine[1:])
    case "to-string": ret = cirruToString(env, codeLine[1:])
    case "string":    ret = cirruString  (env, codeLine[1:])
    case "get":       ret = cirruGet     (env, codeLine[1:])
    case "set":       ret = cirruSet     (env, codeLine[1:])
    case "int":       ret = cirruInt     (env, codeLine[1:])
    case "print":     ret = cirruPrint   (env, codeLine[1:])
    case "bool":      ret = cirruBool    (env, codeLine[1:])
    }
    return
  }
  if headExpression, ok := head.(cirru.List); ok {
    debugPrint(headExpression)
    return
  }
  return
}