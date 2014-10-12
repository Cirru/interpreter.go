
package interpreter

import (
  "strings"
)

func (env *scope) _string(xs sequence) (ret unitype) {
  pieces := []string{}
  for _, piece := range xs {
    key := env.getSymbol(piece)
    pieces = append(pieces, key)
  }
  ret.Type = uniString
  ret.Value = strings.Join(pieces, " ")
  return
}
