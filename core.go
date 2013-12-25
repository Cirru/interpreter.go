
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
  "strings"
  "fmt"
  "strconv"
)

func cirruEcho(env *Env, xs []interface{}) (ret cirruObject) {
  fmt.Println(cirruToString(env, xs).Value)
  return
}

func cirruToString(env *Env, xs []interface{}) (ret cirruObject) {
  hold := []string{}
  for _, item := range xs {
    if buffer, ok := item.(cirru.BufferObj); ok {
      hold = append(hold, buffer.Text)
    }
    if bufferList, ok := item.([]interface{}); ok {
      tmp := cirruToString(env, bufferList).Value
      if tmpString, ok := tmp.(string); ok {
        hold = append(hold, "(" + tmpString + ")")
      }
    }
  }
  ret.Typing = "string"
  ret.Value = strings.Join(hold, " ")
  return
}

func cirruInt(env *Env, xs []interface{}) (ret cirruObject) {
  if textBuffer, ok := xs[0].(cirru.BufferObj); ok {
    intNumber, err := strconv.Atoi(textBuffer.Text)
    if err != nil {
      panic(err)
    }
    ret.Typing = "int"
    ret.Value = intNumber
  }
  return
}

func cirruString(env *Env, xs []interface{}) (ret cirruObject) {
  if textBuffer, ok := xs[0].(cirru.BufferObj); ok {
    ret.Typing = "string"
    ret.Value = textBuffer.Text
  }
  if bufferList, ok := xs[0].([]interface{}); ok {
    ret.Typing = "string"
    ret.Value = cirruToString(env, bufferList).Value
  }
  return
}

func cirruPrint(env *Env, xs []interface{}) (ret cirruObject) {
  outList := []string{}
  for _, value := range xs {
    if textBuffer, ok := value.(cirru.BufferObj); ok {
      list := []interface{}{}
      list = append(list, textBuffer)
      unit := stringifyObject(cirruGet(env, list))
      outList = append(outList, unit)
    }
    if bufferList, ok := value.([]interface{}); ok {
      calculated := Evaluate(env, bufferList)
      // fmt.Println("value is:", calculated)
      unit := stringifyObject(calculated)
      outList = append(outList, unit)
    }
  }
  fmt.Println(strings.Join(outList, " "))
  return
}

func cirruGet(env *Env, xs []interface{}) (ret cirruObject) {

  return
}
