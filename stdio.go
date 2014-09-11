
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
  for _, value := range xs {
    if t, ok := value.(token); ok {
      seq := sequence{}
      seq = append(seq, t)
      piece := stringifyUnitype(env.get(seq))
      outList = append(outList, piece)
    } else if seq, ok := value.(sequence); ok {
      piece := stringifyUnitype(Evaluate(env, seq))
      outList = append(outList, piece)
    }
  }
  fmt.Println(strings.Join(outList, ""))
  return
}
