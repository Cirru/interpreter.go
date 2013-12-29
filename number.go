
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
  "strconv"
)

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
