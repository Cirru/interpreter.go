
package interpreter

type cirruType int

const (
  cirruInt cirruType = iota
  cirruFloat
  cirruBool
  cirruString
  cirruRegexp
  cirruTable
  cirruArray
  cirruFn
  cirruNil
)