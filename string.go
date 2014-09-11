
package interpreter

import (
  "strings"
)

func (env *scope) _string(xs sequence) (ret unitype) {
  pieces := []string{}
  for _, piece := range xs {
    key := env.getKey(piece)
    if key.Type != uniString {
      panic("not a piece of string")
    }
    str, _ := key.Value.(string)
    pieces = append(pieces, str)
  }
  ret.Type = uniString
  ret.Value = strings.Join(pieces, " ")
  return
}
