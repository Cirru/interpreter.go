
package cirruGopher

import (
  "github.com/Cirru/cirru-grammar"
  "fmt"
  "strings"
)

func cirruEcho(env *Env, xs cirru.List) (ret Object) {
  fmt.Println(codeString(xs, 0))
  return
}

func cirruComment(env *Env, xs cirru.List) (ret Object) {
  return
}

func cirruPrint(env *Env, xs cirru.List) (ret Object) {
  outList := []string{}
  for _, value := range xs {
    if token, ok := value.(cirru.Token); ok {
      list := cirru.List{}
      list = append(list, token)
      unit := stringifyObject(cirruGet(env, list), 0)
      outList = append(outList, unit)
    }
    if list, ok := value.(cirru.List); ok {
      calculated := Evaluate(env, list)
      // fmt.Println("value is:", calculated)
      unit := stringifyObject(calculated, 0)
      outList = append(outList, unit)
    }
  }
  fmt.Println(strings.Join(outList, " "))
  return
}
