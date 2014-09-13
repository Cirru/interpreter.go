
package interpreter

import (
  "fmt"
  "strings"
)

func (env *scope) comment(xs sequence) (ret unitype) {
  return uni(nil)
}

func (env *scope) _print(xs sequence) (ret unitype) {
  outList := []string{}
  for _, item := range xs {
    piece := stringifyUnitype(env.getValue(item))
    outList = append(outList, piece)
  }
  fmt.Println(strings.Join(outList, ""))
  return uni(nil)
}
