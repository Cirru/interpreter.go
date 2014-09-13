
package interpreter

import "fmt"

func assertLen(xs []interface{}, n int) {
  if len(xs) != n {
    panic(fmt.Sprintf("length not equal to %d", n))
  }
}

func assert(x bool, msg string) {
  if !x {
    panic(msg)
  }
}