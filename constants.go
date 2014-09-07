
package interpreter

type cirruName int

const (
  cirruInt cirruName = iota
  cirruFloat
  cirruBool
  cirruString
  cirruRegexp
  cirruTable
  cirruArray
  cirruFunction
  cirruCode
  cirruNil
)