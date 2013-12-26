
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
  "strings"
  "fmt"
  "strconv"
)

func cirruEcho(env *Env, xs cirru.List) (ret Object) {
  fmt.Println(cirruToString(env, xs).Value)
  return
}

func cirruToString(env *Env, xs cirru.List) (ret Object) {
  hold := []string{}
  for _, item := range xs {
    if buffer, ok := item.(cirru.Token); ok {
      hold = append(hold, buffer.Text)
    }
    if list, ok := item.(cirru.List); ok {
      tmp := cirruToString(env, list).Value
      if tmpString, ok := tmp.(string); ok {
        hold = append(hold, "(" + tmpString + ")")
      }
    }
  }
  ret.Tag = "string"
  ret.Value = strings.Join(hold, " ")
  return
}

func cirruInt(env *Env, xs cirru.List) (ret Object) {
  if token, ok := xs[0].(cirru.Token); ok {
    intNumber, err := strconv.Atoi(token.Text)
    if err != nil {
      panic(err)
    }
    ret.Tag = "int"
    ret.Value = intNumber
  }
  return
}

func cirruString(env *Env, xs cirru.List) (ret Object) {
  if token, ok := xs[0].(cirru.Token); ok {
    ret.Tag = "string"
    ret.Value = token.Text
  }
  if list, ok := xs[0].(cirru.List); ok {
    ret.Tag = "string"
    ret.Value = cirruToString(env, list).Value
  }
  return
}

func cirruPrint(env *Env, xs cirru.List) (ret Object) {
  outList := []string{}
  for _, value := range xs {
    if token, ok := value.(cirru.Token); ok {
      list := cirru.List{}
      list = append(list, token)
      unit := stringifyObject(cirruGet(env, list))
      outList = append(outList, unit)
    }
    if list, ok := value.(cirru.List); ok {
      calculated := Evaluate(env, list)
      // fmt.Println("value is:", calculated)
      unit := stringifyObject(calculated)
      outList = append(outList, unit)
    }
  }
  fmt.Println(strings.Join(outList, " "))
  return
}

func cirruSet(env *Env, xs cirru.List) (ret Object) {
  value := cirruGet(env, xs[1:2])
  if token, ok := xs[0].(cirru.Token); ok {
    (*env)[token.Text] = value
    return value
  }
  if list, ok := xs[0].(cirru.List); ok {
    variable := cirruGet(env, list[0:1])
    if variable.Tag == "string" {
      if name, ok := variable.Value.(string); ok {
        (*env)[name] = value
        return value
      }
    }
  }
  return
}

func cirruGet(env *Env, xs cirru.List) (ret Object) {
  if token, ok := xs[0].(cirru.Token); ok {
    ret = (*env)[token.Text]
    println("get...", ret.Tag)
    return
  }
  if list, ok := xs[0].(cirru.List); ok {
    ret = Evaluate(env, list)
    return
  }
  return
}