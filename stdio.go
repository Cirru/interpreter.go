
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
  "fmt"
  "strings"
)

func cirruEcho(env *Env, xs cirru.List) (ret Object) {
  fmt.Println(cirruToString(env, xs).Value)
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
