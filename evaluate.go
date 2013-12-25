
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
)

type cirruObject struct {
  Typing string
  Value interface{}
}

type Env map[string]interface{}

// Evaluate read expressions and return a single result
func Evaluate(codeLine []interface{}, env *Env) (ret cirruObject) {
  // debugPrint(codeLine, env)
  if len(codeLine) == 0 {
    emptyArray := cirruObject{"list", codeLine}
    return emptyArray
  }

  head := codeLine[0]
  body := []interface{}{}
  for _, value := range codeLine[1:] {
    body = append(body, value)
  }

  if headBuffer, ok := head.(cirru.BufferObj); ok {
    // debugPrint(headBuffer.Text)
    switch headBuffer.Text {
    case "echo":
      cirruEcho(body...)
    }
    return
  }
  if headExpression, ok := head.([]interface{}); ok {
    debugPrint(headExpression)
    return
  }
  return
}