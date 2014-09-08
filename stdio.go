
package interpreter

import (
  "fmt"
  "strings"
)

func (env *scope) comment(xs sequence) (ret unitype) {
  return unitype{uniNil, nil}
}

func (env *scope) _print(xs sequence) (ret unitype) {
  outList := []string{}
  for _, value := range xs {
    if t, ok := value.(token); ok {
      list := sequence{}
      list = append(list, t)
      unit := stringifyUnitype(env.get(list))
      outList = append(outList, unit)
    } else if list, ok := value.(sequence); ok {
      calculated := Evaluate(env, list)
      // fmt.Println("value is:", calculated)
      unit := stringifyUnitype(calculated)
      outList = append(outList, unit)
    }
  }
  fmt.Println(strings.Join(outList, ""))
  return
}
