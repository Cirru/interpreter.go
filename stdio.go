
package interpreter

import (
  "github.com/Cirru/parser"
  "fmt"
  "strings"
)

func (env *Env) comment(xs []interface{}) (ret unitype) {
  return
}

func (env *Env) _print(xs []interface{}) (ret unitype) {
  outList := []string{}
  for _, value := range xs {
    if token, ok := value.(parser.Token); ok {
      list := []interface{}{}
      list = append(list, token)
      unit := stringifyunitype(env.get(list))
      outList = append(outList, unit)
    }
    if list, ok := value.([]interface{}); ok {
      calculated := Evaluate(env, list)
      // fmt.Println("value is:", calculated)
      unit := stringifyunitype(calculated)
      outList = append(outList, unit)
    }
  }
  fmt.Println(strings.Join(outList, "\n"))
  return
}
