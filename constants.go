
package interpreter

type unitypeName int

const (
  uniInt unitypeName = iota
  uniFloat
  uniBool
  uniString
  uniRegexp
  uniTable
  uniArray
  uniFn
  uniNil
)