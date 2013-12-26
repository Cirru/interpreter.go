
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
)

type cirruObject struct {
  Typing string
  Value interface{}
}

type Env map[string]cirruObject

// Evaluate read expressions and return a single result
func Evaluate(env *Env, codeLine cirru.List) (ret cirruObject) {
  // debugPrint(codeLine, env)
  if len(codeLine) == 0 {
    emptyArray := cirruObject{"list", codeLine}
    return emptyArray
  }

  head := codeLine[0]

  if headBuffer, ok := head.(cirru.Token); ok {
    // debugPrint(headBuffer.Text)
    switch headBuffer.Text {
    case "echo":      ret = cirruEcho    (env, codeLine[1:])
    case "to-string": ret = cirruToString(env, codeLine[1:])
    case "get":       ret = cirruGet     (env, codeLine[1:])
    case "int":       ret = cirruInt     (env, codeLine[1:])
    case "print":     ret = cirruPrint   (env, codeLine[1:])
    }
    return
  }
  if headExpression, ok := head.(cirru.List); ok {
    debugPrint(headExpression)
    return
  }
  return
}