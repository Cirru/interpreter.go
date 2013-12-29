
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
  // debugPrint("cirruSet:", xs)
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
    if value, ok := (*env)[token.Text]; ok {
      ret = value
      return
    } else {
      if parent, ok := (*env)["parent"]; ok {
        if parent.Tag == "map" {
          if scope, ok := parent.Value.(*Env); ok {
            ret = cirruGet(scope, xs[0:1])
            return
          }
        }
      }
    }
    return
  }
  if list, ok := xs[0].(cirru.List); ok {
    ret = Evaluate(env, list)
    return
  }
  return
}

func cirruBool(env *Env, xs cirru.List) (ret Object) {
  ret.Tag = "bool"
  ret.Value = false
  if token, ok := xs[0].(cirru.Token); ok {
    trueValues := []string{"true", "yes", "riight", "1"}
    for _, text := range trueValues {
      if text == token.Text {
        ret.Value = true
      }
    }
    return
  }
  return
}

func cirruFloat(env *Env, xs cirru.List) (ret Object) {
  if token, ok := xs[0].(cirru.Token); ok {
    floatNumber, err := strconv.ParseFloat(token.Text, 64)
    if err != nil {
      panic(err)
    }
    ret.Tag = "float"
    ret.Value = floatNumber
  }
  return
}

func cirruType(env *Env, xs cirru.List) (ret Object) {
  value := cirruGet(env, xs[0:1])
  if &value != nil {
    ret.Tag = "string"
    ret.Value = value.Tag
  }
  return
}

func cirruArray(env *Env, xs cirru.List) (ret Object) {
  ret.Tag = "array"
  hold := []Object{}
  for _, item := range xs {
    list := cirru.List{item}
    hold = append(hold, cirruGet(env, list))
  }
  tmp := []interface{}{}
  tmp = append(tmp, &hold)
  ret.Value = tmp[0]
  return
}

func cirruMap(env *Env, xs cirru.List) (ret Object) {
  ret.Tag = "map"
  hold := Env{}
  for _, item := range xs {
    if pair, ok := item.(cirru.List); ok {
      name := pair[0]
      var key string
      if token, ok := name.(cirru.Token); ok {
        key = token.Text
      }
      value := cirruGet(env, pair[1:2])
      hold[key] = value
    }
  }
  ret.Value = &hold
  return
}

func cirruSelf(env *Env, xs cirru.List) (ret Object) {
  ret.Tag = "map"
  ret.Value = env
  return
}

func cirruChild(env *Env, xs cirru.List) (ret Object) {
  childMap := Env{}
  childMap["parent"] = cirruSelf(env, xs)
  ret.Tag = "map"
  ret.Value = &childMap
  println("ret is:", ret.Value)
  return
}

func cirruUnder(env *Env, xs cirru.List) (ret Object) {
  item := cirruGet(env, xs[0:1])
  // debugPrint("item is:", item.Value)
  if scope, ok := item.Value.(*Env); ok {
    debugPrint("scope is:", xs[1])
    for _, exp := range xs[1:] {
      if list, ok := exp.(cirru.List); ok {
        ret = Evaluate(scope, list)
      }
    }
  } else {
    debugPrint("no scope", item.Value, xs[1:])
  }
  return
}