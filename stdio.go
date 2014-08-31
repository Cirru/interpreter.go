
package interpreter

import (
  "github.com/Cirru/parser"
  "fmt"
  "strings"
)

func cirruComment(env *Env, xs []interface{}) (ret Object) {
  return
}

func cirruPrint(env *Env, xs []interface{}) (ret Object) {
  outList := []string{}
  for _, value := range xs {
    if token, ok := value.(parser.Token); ok {
      list := []interface{}{}
      list = append(list, token)
      unit := stringifyObject(cirruGet(env, list), 0)
      outList = append(outList, unit)
    }
    if list, ok := value.([]interface{}); ok {
      calculated := Evaluate(env, list)
      // fmt.Println("value is:", calculated)
      unit := stringifyObject(calculated, 0)
      outList = append(outList, unit)
    }
  }
  fmt.Println(strings.Join(outList, "\n"))
  return
}
