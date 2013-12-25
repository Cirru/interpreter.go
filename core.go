
package cirruGopher

import (
  "github.com/jiyinyiyong/cirru-grammar"
  "strings"
  "fmt"
)

func cirruEcho(xs ...interface{}) {
  fmt.Println(cirruToString(xs))
}

func cirruToString(xs []interface{}) (ret string) {
  hold := []string{}
  for _, item := range xs {
    if buffer, ok := item.(cirru.BufferObj); ok {
      hold = append(hold, buffer.Text)
    }
    if bufferList, ok := item.([]interface{}); ok {
      tmp := cirruToString(bufferList)
      hold = append(hold, "(" + tmp + ")")
    }
  }
  ret = strings.Join(hold, " ")
  return
}